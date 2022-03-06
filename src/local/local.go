package local

import (
	"sync"
)

type localMessenger struct {
	sendChannels    map[string]chan []byte
	receiveChannels map[string]map[string]chan []byte

	wg sync.WaitGroup
	sync.Mutex

	errors         chan error
	doneSender     chan struct{}
	doneReveciever chan struct{}
}

func NewLocalMessenger() *localMessenger {
	return &localMessenger{
		receiveChannels: make(map[string]map[string]chan []byte),
		sendChannels:    make(map[string]chan []byte),

		errors:         make(chan error),
		doneSender:     make(chan struct{}),
		doneReveciever: make(chan struct{}),
	}
}

func (m *localMessenger) Receive(exchange string, name string) <-chan []byte {
	m.Lock()
	defer m.Unlock()

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

func (m *localMessenger) Errors() <-chan error {
	return m.errors
}

func (m *localMessenger) Stop() {
	close(m.doneSender)
	close(m.doneReveciever)
	m.wg.Wait()
}

func (m *localMessenger) newReceiver(exchange string, name string) (chan []byte, error) {
	receiver := make(chan []byte, 1024)
	m.wg.Add(1)
	defer m.wg.Done()
	return receiver, nil
}

func (m *localMessenger) newSender(exchange string) (chan []byte, error) {
	sender := make(chan []byte, 1024)
	m.wg.Add(1)
	go func() {
		defer m.wg.Done()
		for {
			select {
			case message := <-sender:
				for _, receiver := range m.receiveChannels[exchange] {
					receiver <- message
				}
			case <-m.doneSender:
				return
			}
		}
	}()
	return sender, nil
}
