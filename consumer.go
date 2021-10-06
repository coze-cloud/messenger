package messenger

type Consumer struct {
	handler handler
	name string
	autoAcknowledge bool
}

func NewConsumer(handler handler) Consumer {
	return Consumer{handler: handler}
}

func (c Consumer) Name(name string) Consumer {
	c.name = name

	return c
}

func (c Consumer) ShouldAutoAcknowledge() Consumer {
	c.autoAcknowledge = true

	return c
}
