package messenger

import "fmt"

type Consumer struct {
	handler handler
	name string
	autoAcknowledge bool

	address *address
}

func NewConsumer(handler handler) Consumer {
	return Consumer{handler: handler}
}

func (c Consumer) Named(name string) Consumer {
	c.name = name

	return c
}

func (c Consumer) ShouldAutoAcknowledge() Consumer {
	c.autoAcknowledge = true

	return c
}

func (c Consumer) locatedAt(address *address) Consumer {
	c.address = address

	name := address.String()
	if len(c.name) > 0 {
		name = fmt.Sprintf("%s@%s", c.name, name)
	}
	c.name = name

	return c
}
