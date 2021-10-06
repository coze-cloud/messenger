package messenger

import "strings"

type createRabbitExchangeCommand struct {
	channel Channel
	exchange Exchange
}

func newCreateRabbitExchangeCommand(channel Channel, exchange Exchange) Command {
	return &createRabbitExchangeCommand{channel: channel, exchange: exchange}
}

func (c createRabbitExchangeCommand) Handle() error {
	exchangeNameLength := len(strings.TrimSpace(c.exchange.name))
	if exchangeNameLength <= 0 {
		// Define no exchange because the default
		// exchange always exists in Rabbit MQ
		return nil
	}

	return c.channel.DeclareExchange(c.exchange)
}

