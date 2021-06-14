package messenger

import (
	"github.com/streadway/amqp"
	"strconv"
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
		args["x-message-ttl"] = strconv.FormatInt(factory.queue.TimeToLive, 10)
	}

	return factory.channel.QueueDeclare(
		factory.queue.Name,
		false,
		factory.queue.ShouldAutoDelete,
		false,
		false,
		args,
	)
}
