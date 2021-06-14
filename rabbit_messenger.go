package messenger

import (
	"github.com/mcuadros/go-defaults"
	"github.com/streadway/amqp"
)

type rabbitMessenger struct {
	Messenger // Interface

	connection *amqp.Connection
	address    address
}

func NewRabbitMessenger(url string) (Messenger, error) {
	messenger := new(rabbitMessenger)
	defaults.SetDefaults(messenger)
	messenger.address = newRandomAddress()

	var err error

	messenger.connection, err = amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	return messenger, nil
}

func (messenger rabbitMessenger) GetAddress() address {
	return messenger.address
}

func (messenger *rabbitMessenger) Publish(queue Queue, publication Publication) error {
	channel, err := messenger.connection.Channel()
	if err != nil {
		return err
	}

	rabbitQueue, err := NewRabbitQueueFactory(channel, queue).Produce()
	if err != nil {
		return err
	}

	err = newRabbitPublisher(messenger.address, channel, publication).Publish(rabbitQueue)
	if err != nil {
		return err
	}

	if err := channel.Close(); err != nil {
		return err
	}
	return nil
}

func (messenger *rabbitMessenger) Consume(queue Queue, consumption Consumption) error {
	channel, err := messenger.connection.Channel()
	if err != nil {
		return err
	}

	rabbitQueue, err := NewRabbitQueueFactory(channel, queue).Produce()
	if err != nil {
		return err
	}

	err = newRabbitConsumer(messenger.address, channel, consumption).Consume(rabbitQueue)
	if err != nil {
		return err
	}

	return nil
}