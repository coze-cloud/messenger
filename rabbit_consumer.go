package messenger

import (
	"fmt"
	"github.com/streadway/amqp"
	"strings"
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

func (consumer rabbitConsumer) Consume(exchange Exchange, queue Queue) error {
	name := consumer.receiver.String()
	if len(strings.TrimSpace(consumer.consumption.Name)) > 0 {
		name = fmt.Sprintf("%s@%s", consumer.consumption.Name, name)
	}

	deliveries, err := consumer.channel.Consume(
		queue.Name,
		name,
		consumer.consumption.IsAutoAcknowledge,
		false,
		false,
		false,
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