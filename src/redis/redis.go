package redis

import (
	"sync"

	"github.com/go-redis/redis"
)

type Options = redis.Options

type redisMessenger struct {
	receiverChannels map[string]chan []byte
	sendChannels     map[string]chan []byte

	wg sync.WaitGroup
	sync.Mutex

	errors         chan error
	doneSender     chan struct{}
	doneReveciever chan struct{}

	options *Options
}

func NewRedisClient(options *Options) *redisMessenger {
	return &redisMessenger{
		receiverChannels: make(map[string]chan []byte),
		sendChannels:     make(map[string]chan []byte),

		errors:         make(chan error),
		doneSender:     make(chan struct{}),
		doneReveciever: make(chan struct{}),

		options: options,
	}
}

func (m *redisMessenger) Receive(topic string) <-chan []byte {
	m.Lock()
	defer m.Unlock()

	receiver, ok := m.sendChannels[topic]
	if ok {
		return receiver
	}

	receiveChannel, err := m.newReceiver(topic)
	if err != nil {
		m.errors <- err
	}
	m.receiverChannels[topic] = receiveChannel

	return receiveChannel
}

func (m *redisMessenger) Send(topic string) chan<- []byte {
	m.Lock()
	defer m.Unlock()

	sendChannel, ok := m.sendChannels[topic]
	if ok {
		return sendChannel
	}

	sendChannel, err := m.newSender(topic)
	if err != nil {
		m.errors <- err
	}
	m.sendChannels[topic] = sendChannel

	return sendChannel
}

func (m *redisMessenger) Errors() <-chan error {
	return m.errors
}

func (m *redisMessenger) Stop() {
	close(m.doneSender)
	close(m.doneReveciever)
	m.wg.Wait()
}

func (m *redisMessenger) newClient() (*redis.Client, error) {
	client := redis.NewClient(m.options)
	if _, err := client.Ping().Result(); err != nil {
		return nil, err
	}
	return client, nil
}

func (m *redisMessenger) newReceiver(topic string) (chan []byte, error) {
	client, err := m.newClient()
	if err != nil {
		return nil, err
	}

	receiver := make(chan []byte, 1024)
	m.wg.Add(1)
	go func() {
		defer m.wg.Done()
		defer client.Close()

		subscriber := client.PSubscribe(topic)
		defer subscriber.Close()
		for {
			select {
			case <-m.doneReveciever:
				return
			case message := <-subscriber.Channel():
				receiver <- []byte(message.Payload)
			}
		}
	}()
	return receiver, nil
}

func (m *redisMessenger) newSender(topic string) (chan []byte, error) {
	client, err := m.newClient()
	if err != nil {
		return nil, err
	}

	sender := make(chan []byte, 1024)
	m.wg.Add(1)
	go func() {
		defer m.wg.Done()
		defer client.Close()

		for {
			select {
			case <-m.doneSender:
				return
			case message := <-sender:
				if client.Publish(topic, message).Err(); err != nil {
					m.errors <- err
				}
			}

		}
	}()
	return sender, nil
}
