package messenger

import (
	"github.com/streadway/amqp"
)

type rabbitQueueFactory struct {
	channel *amqp.Channel
	queue   Queue
}

func newRabbitQueueFactory(channel *amqp.Channel, queue Queue) rabbitQueueFactory {
	producer := new(rabbitQueueFactory)
	producer.channel = channel
	producer.queue = queue

	return *producer
}

func (factory rabbitQueueFactory) Produce() (amqp.Queue, error) {
	args := make(map[string]interface{})
	if factory.queue.TimeToLive > 0 {
		args["x-message-ttl"] = factory.queue.TimeToLive
	}

	return factory.channel.QueueDeclare(
		factory.queue.Name,
		false,
		factory.queue.ShouldAutoRemove || len(factory.queue.Name) == 0,
		false,
		false,
		args,
	)
}
