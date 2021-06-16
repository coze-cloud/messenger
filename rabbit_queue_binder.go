package messenger

import (
	"github.com/streadway/amqp"
	"strings"
)

type rabbitQueueBinder struct {
	channel *amqp.Channel
	exchange Exchange
	queue Queue
}

func newRabbitQueueBinder(channel *amqp.Channel, exchange Exchange, queue Queue) rabbitQueueBinder {
	binder := new(rabbitQueueBinder)
	binder.channel = channel
	binder.exchange = exchange
	binder.queue = queue

	return *binder
}

func (binder rabbitQueueBinder) Bind() error {
	if len(strings.TrimSpace(binder.exchange.Name)) <= 0 {
		return nil
	}

	if err := binder.channel.QueueBind(
		binder.queue.Name,
		binder.queue.Topic,
		binder.exchange.Name,
		false,
		nil,
	); err != nil {
		return err
	}

	return nil
}