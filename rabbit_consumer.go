package messenger

import (
	"github.com/streadway/amqp"
)

type rabbitConsumer struct {
	receiver    address
	channel     *amqp.Channel
	consumption Consumption
}

func newRabbitConsumer(receiver address, channel *amqp.Channel, consumption Consumption) rabbitConsumer {
	consumer := new(rabbitConsumer)
	consumer.receiver = receiver
	consumer.channel = channel
	consumer.consumption = consumption

	return *consumer
}

func (consumer rabbitConsumer) Consume(queue amqp.Queue) error {
	deliveries, err := consumer.channel.Consume(
		queue.Name,
		consumer.consumption.Name,
		consumer.consumption.IsAutoAcknowledge,
		consumer.consumption.IsExclusive,
		consumer.consumption.IsNoLocal,
		consumer.consumption.IsNoWait,
		nil,
	)
	if err != nil {
		return err
	}

	go consumer.deliver(deliveries)
	return nil
}

func (consumer rabbitConsumer) deliver(deliveries <-chan amqp.Delivery) {
	for delivery := range deliveries {
		context := newRabbitHandlerContext(consumer.receiver, consumer.channel, delivery)
		consumer.consumption.Handler(context)
	}
}