package messenger

import "github.com/streadway/amqp"

type rabbitConnection struct {
	connection *amqp.Connection
}

func newRabbitConnection(url string) (*rabbitConnection, error) {
	connection, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}
	return &rabbitConnection{connection: connection}, nil
}

func (r rabbitConnection) GetChannel() (Channel, error) {
	rabbitChannel, err := r.connection.Channel()
	if err != nil {
		return nil, err
	}
	return newRabbitChannel(rabbitChannel), nil
}

func (r rabbitConnection) Close() error {
	return r.connection.Close()
}

