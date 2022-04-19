package redis

import (
	"context"
	"fmt"
	"sync"

	"github.com/go-redis/redis/v8"
	uuid "github.com/satori/go.uuid"
)

type redisStreamMessenger struct {
	receiveChannels map[string]chan []byte
	sendChannels    map[string]chan []byte

	mutex sync.Mutex

	errors        chan error
	receiveCtx    context.Context
	receiveCancel context.CancelFunc
	sendCtx       context.Context
	sendCancel    context.CancelFunc

	address string
}

func NewRedisStreamMessenger(ctx context.Context, address string) *redisStreamMessenger {
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
	}
}

func (m *redisStreamMessenger) Receive(exchange string, name string) <-chan []byte {
	m.mutex.Lock()
	defer m.mutex.Unlock()

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
	m.mutex.Lock()
	defer m.mutex.Unlock()

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

	receiver := make(chan []byte)
	go func() {
		connection.XGroupCreateMkStream(m.receiveCtx, exchange, name, "$")

		for {
			select {
			case <-m.receiveCtx.Done():
				close(receiver)
				return
			default:
				streams, err := connection.XReadGroup(m.receiveCtx, &redis.XReadGroupArgs{
					Streams:  []string{exchange, ">"},
					Group:    name,
					Consumer: uuid.NewV4().String(),
				}).Result()
				if err != nil {
					switch err {
					case redis.Nil:
						continue
					case context.Canceled:
						continue
					case context.DeadlineExceeded:
						continue
					default:
						m.errors <- err
						continue
					}
				}

				for _, stream := range streams {
					if stream.Stream != exchange {
						continue
					}

					for _, message := range stream.Messages {
						data, ok := message.Values["message"].(string)
						if !ok {
							m.errors <- fmt.Errorf("invalid message received: %v", message.Values)
							continue
						}
						receiver <- []byte(data)
					}
				}
			}
		}
	}()
	return receiver, nil
}

func (m *redisStreamMessenger) newSender(exchange string) (chan []byte, error) {
	connection, err := m.newConnection()
	if err != nil {
		return nil, err
	}

	sender := make(chan []byte, 1024)
	go func() {
		for {
			select {
			case message := <-sender:
				_, err := connection.XAdd(m.sendCtx, &redis.XAddArgs{
					Stream: exchange,
					ID:     "*",
					Values: map[string]interface{}{
						"message": string(message),
					},
				}).Result()
				if err != nil {
					switch err {
					case context.Canceled:
						continue
					case context.DeadlineExceeded:
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
