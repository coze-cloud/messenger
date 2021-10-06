package messenger

import "github.com/streadway/amqp"

type rabbitChannel struct {
	channel *amqp.Channel
}

func newRabbitChannel(channel *amqp.Channel) *rabbitChannel {
	return &rabbitChannel{channel: channel}
}

func (c rabbitChannel) DeclareExchange(exchange Exchange) error {
	return c.channel.ExchangeDeclare(
		exchange.name,
		exchange.strategy,
		true,
		exchange.autoRemove,
		false,
		false,
		nil,
	)
}

func (c rabbitChannel) DeclareQueue(queue Queue, args map[string]interface{}) error {
	_, err := c.channel.QueueDeclare(
		queue.name,
		false,
		queue.autoRemove,
		false,
		false,
		args,
	)
	return err
}

func (c rabbitChannel) DeleteQueue(queue Queue) error {
	_, err := c.channel.QueueDelete(
		queue.name,
		true,
		false,
		false,
	);
	return err
}

func (c rabbitChannel) BindQueueToExchange(exchange Exchange, queue Queue) error {
	return c.channel.QueueBind(
		queue.name,
		queue.topic,
		exchange.name,
		false,
		nil,
	);
}

func (c rabbitChannel) Publish(exchange Exchange, queue Queue, message Message) error {
	serializer := newJsonMessageSerializer(message)

	body, err := serializer.Serialize()
	if err != nil {
		return err
	}

	return c.channel.Publish(
		exchange.name,
		queue.topic,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/json",
			Body: []byte(body),
		},
	)
}

func (c rabbitChannel) Consume(queue Queue, consumer Consumer) error {
	deliveries, err := c.channel.Consume(
		queue.name,
		consumer.name,
		consumer.autoAcknowledge,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	go func() {
		for range deliveries {
			var context Context // TODO: Construct context
			consumer.handler(context)
		}
	}()

	return nil
}