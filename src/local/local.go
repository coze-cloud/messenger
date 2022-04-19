package local

import (
	"context"
	"sync"
)

type localMessenger struct {
	sendChannels    map[string]chan []byte
	receiveChannels map[string]map[string]chan []byte

	mutex sync.Mutex

	errors        chan error
	receiveCtx    context.Context
	receiveCancel context.CancelFunc
	sendCtx       context.Context
	sendCancel    context.CancelFunc
}

func NewLocalMessenger(ctx context.Context) *localMessenger {
	receiveCtx, receiveCancel := context.WithCancel(ctx)
	sendCtx, sendCancel := context.WithCancel(ctx)

	return &localMessenger{
		receiveChannels: make(map[string]map[string]chan []byte),
		sendChannels:    make(map[string]chan []byte),

		errors:        make(chan error),
		receiveCtx:    receiveCtx,
		receiveCancel: receiveCancel,
		sendCtx:       sendCtx,
		sendCancel:    sendCancel,
	}
}

func (m *localMessenger) Receive(exchange string, name string) <-chan []byte {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if m.receiveChannels[exchange] == nil {
		m.receiveChannels[exchange] = make(map[string]chan []byte)
	}

	receiveChannel, ok := m.receiveChannels[exchange][name]
	if ok {
		return receiveChannel
	}

	receiveChannel, err := m.newReceiver(exchange, name)
	if err != nil {
		m.errors <- err
	}
	m.receiveChannels[exchange][name] = receiveChannel

	return receiveChannel
}

func (m *localMessenger) Send(exchange string) chan<- []byte {
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

func (m *localMessenger) Errors() <-chan error {
	return m.errors
}

func (m *localMessenger) newReceiver(exchange string, name string) (chan []byte, error) {
	receiver := make(chan []byte, 1024)
	go func() {
		<-m.receiveCtx.Done()
		close(receiver)
	}()
	return receiver, nil
}

func (m *localMessenger) newSender(exchange string) (chan []byte, error) {
	sender := make(chan []byte, 1024)
	go func() {
		for {
			select {
			case message := <-sender:
				for _, receiver := range m.receiveChannels[exchange] {
					receiver <- message
				}
			case <-m.sendCtx.Done():
				return
			}
		}
	}()
	return sender, nil
}
