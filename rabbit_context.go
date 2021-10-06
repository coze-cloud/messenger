package messenger

type rabbitContext struct {
	channel Channel
	exchange Exchange
	queue Queue
	delivery Delivery
}

func newRabbitContext(channel Channel, exchange Exchange, queue Queue, delivery Delivery) Context {
	return &rabbitContext{
		channel: channel,
		exchange: exchange,
		queue: queue,
		delivery: delivery,
	}
}

func (c rabbitContext) GetDelivery() Delivery {
	return c.delivery
}


func (c rabbitContext) Publish(message Message) error {
	return c.channel.Publish(c.exchange, c.queue, message)
}

