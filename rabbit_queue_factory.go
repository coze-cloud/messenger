package messenger

import (
	"github.com/streadway/amqp"
)

type rabbitQueueFactory struct {
	channel *amqp.Channel
	queue   Queue
}

func NewRabbitQueueFactory(channel *amqp.Channel, queue Queue) rabbitQueueFactory {
	producer := new(rabbitQueueFactory)
	producer.channel = channel
	producer.queue = queue

	return *producer
}

func (factory rabbitQueueFactory) Produce() (amqp.Queue, error) {
	return factory.channel.QueueDeclare(
		factory.queue.Name,
		factory.queue.IsDurable,
		factory.queue.ShouldDeleteIfUnused,
		factory.queue.IsExclusive,
		factory.queue.IsNoWait,
		nil,
	)
}
