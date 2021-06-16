package messenger

import "github.com/mcuadros/go-defaults"

type Exchange struct {
	Name             string
	Strategy         string `default:"direct"`
	ShouldAutoRemove bool
}

func NewExchange() Exchange {
	exchange := new(Exchange)
	defaults.SetDefaults(exchange)

	return *exchange
}

func (exchange Exchange) Named(name string) Exchange {
	exchange.Name = name
	return exchange
}

func (exchange Exchange) WithStrategy(strategy string) Exchange {
	exchange.Strategy = strategy
	return exchange
}

func (exchange Exchange) AutoRemove() Exchange {
	exchange.ShouldAutoRemove = true
	return exchange
}