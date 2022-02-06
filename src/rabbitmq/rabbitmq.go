package rabbitmq

import (
	"sync"

	"github.com/streadway/amqp"
)

type rabbitmqMessenger struct {
	receiveChannels map[string]chan []byte
	sendChannels    map[string]chan []byte

	wg sync.WaitGroup
	sync.Mutex

	errors         chan error
	doneSender     chan struct{}
	doneReveciever chan struct{}

	address     string
	connenction *amqp.Connection
}

func NewRabbitMesseger(address string) *rabbitmqMessenger {
	return &rabbitmqMessenger{
		sendChannels:    make(map[string]chan []byte),
		receiveChannels: make(map[string]chan []byte),

		errors:         make(chan error),
		doneSender:     make(chan struct{}),
		doneReveciever: make(chan struct{}),

		address:     address,
		connenction: nil,
	}
}

func (m *rabbitmqMessenger) Receive(exchange string, name string) <-chan []byte {
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

func (m *rabbitmqMessenger) Send(exchange string) chan<- []byte {
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

func (m *rabbitmqMessenger) Errors() <-chan error {
	return m.errors
}

func (m *rabbitmqMessenger) Stop() {
	close(m.doneSender)
	close(m.doneReveciever)
	m.wg.Wait()
}

func (m *rabbitmqMessenger) newConnection() (*amqp.Connection, error) {
	connection, err := amqp.Dial(m.address)
	if err != nil {
		return nil, err
	}
	return connection, nil
}

func (m *rabbitmqMessenger) newReceiver(exchange string, name string) (chan []byte, error) {
	connection, err := m.newConnection()
	if err != nil {
		return nil, err
	}

	channel, err := connection.Channel()
	if err != nil {
		return nil, err
	}

	if err := channel.ExchangeDeclare(
		exchange,
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		return nil, err
	}

	if _, err := channel.QueueDeclare(
		name,
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		return nil, err
	}

	if err = channel.QueueBind(
		name,
		"",
		exchange,
		false,
		nil,
	); err != nil {
		return nil, err
	}

	messages, err := channel.Consume(
		name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	receiver := make(chan []byte, 1024)
	m.wg.Add(1)
	go func() {
		defer m.wg.Done()
		defer channel.Close()

		for {
			select {
			case message := <-messages:
				receiver <- message.Body
				message.Ack(false)
			case <-m.doneReveciever:
				return
			}
		}
	}()
	return receiver, nil
}

func (m *rabbitmqMessenger) newSender(exchange string) (chan []byte, error) {
	connection, err := m.newConnection()
	if err != nil {
		return nil, err
	}

	channel, err := connection.Channel()
	if err != nil {
		return nil, err
	}

	if err := channel.ExchangeDeclare(
		exchange,
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		return nil, err
	}

	sender := make(chan []byte, 1024)
	m.wg.Add(1)
	go func() {
		defer m.wg.Done()
		defer channel.Close()

		for {
			select {
			case message := <-sender:
				if err := channel.Publish(
					exchange,
					"",
					false,
					false,
					amqp.Publishing{
						Body: message,
					},
				); err != nil {
					m.errors <- err
				}
			case <-m.doneSender:
				return
			}
		}
	}()
	return sender, nil
}
