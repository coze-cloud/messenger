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

func (messenger *rabbitMessenger) Publish(exchange Exchange, queue Queue, publication Publication) error {
	channel, err := messenger.connection.Channel()
	if err != nil {
		return err
	}

	if err = newRabbitExchangeFactory(channel, exchange).Produce(); err != nil {
		return err
	}

	if _, err = newRabbitQueueFactory(channel, queue).Produce(); err != nil {
		return err
	}

	if err = newRabbitQueueBinder(channel, exchange, queue).Bind(); err != nil {
		return err
	}

	if err = newRabbitPublisher(messenger.address, channel, publication).
		Publish(exchange, queue); err != nil {
		return err
	}

	// Delete queue if it is an auto-generated one
	if len(queue.Name) == 0 {
		if _, err := channel.QueueDelete(queue.Name, true, false, false); err != nil {
			return err
		}
	}

	if err := channel.Close(); err != nil {
		return err
	}
	return nil
}

func (messenger *rabbitMessenger) Consume(exchange Exchange,  queue Queue, consumption Consumption) (func() error, error) {
	channel, err := messenger.connection.Channel()
	channelClose := func() error {
		return channel.Close()
	}

	if err != nil {
		return channelClose, err
	}

	if err = newRabbitExchangeFactory(channel, exchange).Produce(); err != nil {
		return channelClose, err
	}

	if _, err = newRabbitQueueFactory(channel, queue).Produce(); err != nil {
		return channelClose, err
	}

	if err = newRabbitQueueBinder(channel, exchange, queue).Bind(); err != nil {
		return channelClose, err
	}

	if err = newRabbitConsumer(messenger.address, channel, consumption).
		Consume(exchange, queue); err != nil {
		return channelClose, err
	}

	return channelClose, nil
}

func (messenger rabbitMessenger) Close(errorHandler func(err error)) {
	err := messenger.connection.Close()
	if err != nil && errorHandler != nil {
		errorHandler(err)
	}
}