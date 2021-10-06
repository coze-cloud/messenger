package messenger

import "strings"

type bindRabbitQueueExchangeCommand struct {
	channel Channel
	exchange Exchange
	queue Queue
}

func newBindRabbitQueueExchangeCommand(channel Channel, exchange Exchange, queue Queue) Command {
	return &bindRabbitQueueExchangeCommand{channel: channel, exchange: exchange, queue: queue}
}

func (c bindRabbitQueueExchangeCommand) Handle() error {
	if len(strings.TrimSpace(c.exchange.name)) <= 0 {
		// Binding on default exchange is not needed
		// and always done automatically
		return nil
	}

	return c.channel.BindQueueToExchange(c.exchange, c.queue)
}
