package redis

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

type redisStreamMessenger struct {
	receiveChannels map[string]chan []byte
	sendChannels    map[string]chan []byte

	wg sync.WaitGroup
	sync.Mutex

	errors        chan error
	receiveCtx    context.Context
	receiveCancel context.CancelFunc
	sendCtx       context.Context
	sendCancel    context.CancelFunc

	address string
	stream  string
}

func NewRedisStreamMessenger(ctx context.Context, address string, stream string) *redisStreamMessenger {
	receiveCtx, receiveCancel := context.WithCancel(ctx)
	sendCtx, sendCancel := context.WithCancel(ctx)

	return &redisStreamMessenger{
		sendChannels:    make(map[string]chan []byte),
		receiveChannels: make(map[string]chan []byte),

		errors:        make(chan error),
		receiveCtx:    receiveCtx,
		receiveCancel: receiveCancel,
		sendCtx:       sendCtx,
		sendCancel:    sendCancel,

		address: address,
		stream:  stream,
	}
}

func (m *redisStreamMessenger) Receive(exchange string, name string) <-chan []byte {
	m.Lock()
	defer m.Unlock()

	receiveChannel, ok := m.receiveChannels[name]
	if ok {
		return receiveChannel
	}

	receiveChannel, err := m.newReceiver(exchange, name)
	if err != nil {
		m.errors <- err
	}
	m.receiveChannels[name] = receiveChannel

	return receiveChannel
}

func (m *redisStreamMessenger) Send(exchange string) chan<- []byte {
	m.Lock()
	defer m.Unlock()

	sendChannel, ok := m.sendChannels[exchange]
	if ok {
		return sendChannel
	}

	sendChannel, err := m.newSender(exchange)
	if err != nil {
		m.errors <- err
	}
	m.sendChannels[exchange] = sendChannel

	return sendChannel
}

func (m *redisStreamMessenger) Errors() <-chan error {
	return m.errors
}

func (m *redisStreamMessenger) Stop() {
	m.sendCancel()
	m.receiveCancel()
	m.wg.Wait()
}

func (m *redisStreamMessenger) newConnection() (*redis.Client, error) {
	opt, err := redis.ParseURL(m.address)
	if err != nil {
		return nil, err
	}
	return redis.NewClient(opt), nil
}

func (m *redisStreamMessenger) newReceiver(exchange string, name string) (chan []byte, error) {
	connection, err := m.newConnection()
	if err != nil {
		return nil, err
	}

	receiveChannel := make(chan []byte)
	m.wg.Add(1)
	go func() {
		defer m.wg.Done()

		connection.XGroupCreateMkStream(m.receiveCtx, m.stream, exchange, "0")

		for {
			select {
			case <-m.receiveCtx.Done():
				return
			default:
				streams, err := connection.XReadGroup(m.receiveCtx, &redis.XReadGroupArgs{
					Streams:  []string{m.stream, ">"},
					Group:    exchange,
					Consumer: name,
					Count:    1,
					Block:    time.Second,
				}).Result()
				if err != nil {
					switch err {
					case redis.Nil:
						continue
					case context.Canceled:
						continue
					default:
						m.errors <- err
						continue
					}
				}

				for _, stream := range streams {
					if stream.Stream != m.stream {
						continue
					}

					for _, message := range stream.Messages {
						data, ok := message.Values["message"].(string)
						if !ok {
							m.errors <- fmt.Errorf("invalid message received: %v", message.Values)
							continue
						}
						receiveChannel <- []byte(data)
					}
				}
			}
		}
	}()
	return receiveChannel, nil
}

func (m *redisStreamMessenger) newSender(exchange string) (chan []byte, error) {
	connection, err := m.newConnection()
	if err != nil {
		return nil, err
	}

	sender := make(chan []byte, 1024)
	m.wg.Add(1)
	go func() {
		defer m.wg.Done()

		connection.XGroupCreateMkStream(m.receiveCtx, m.stream, exchange, "0")

		for {
			select {
			case message := <-sender:
				_, err := connection.XAdd(m.sendCtx, &redis.XAddArgs{
					Stream: m.stream,
					ID:     "*",
					Values: map[string]interface{}{
						"message": string(message),
					},
				}).Result()
				if err != nil {
					switch err {
					case context.Canceled:
						continue
					default:
						m.errors <- err
						continue
					}
				}
			case <-m.sendCtx.Done():
				return
			}
		}
	}()
	return sender, nil
}
