package messenger

import (
	"github.com/streadway/amqp"
	"strings"
)

type rabbitExchangeFactory struct {
	channel  *amqp.Channel
	exchange Exchange
}

func newRabbitExchangeFactory(channel *amqp.Channel, exchange Exchange) rabbitExchangeFactory {
	factory := new(rabbitExchangeFactory)
	factory.channel = channel
	factory.exchange = exchange

	return *factory
}

func (factory rabbitExchangeFactory) Produce() error {
	if len(strings.TrimSpace(factory.exchange.Name)) == 0 {
		return nil
	}

	return factory.channel.ExchangeDeclare(
		factory.exchange.Name,
		factory.exchange.Strategy,
		true,
		factory.exchange.ShouldAutoRemove,
		false,
		false,
		nil,
	)
}