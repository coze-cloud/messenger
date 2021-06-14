package messenger

import "github.com/mcuadros/go-defaults"

type Consumption struct {
	Handler           handler
	Name              string `default:""`
	IsAutoAcknowledge bool `default:"false"`
	IsExclusive       bool `default:"false"`
	IsNoLocal         bool `default:"false"`
	IsNoWait          bool `default:"false"`
}

func NewConsumption(handler handler) Consumption {
	consumption := new(Consumption)
	defaults.SetDefaults(consumption)
	consumption.Handler = handler

	return *consumption
}

func (consumption Consumption) Named(name string) Consumption {
	consumption.Name = name
	return consumption
}

func (consumption Consumption) AutoAcknowledge() Consumption {
	consumption.IsAutoAcknowledge = true
	return consumption
}

func (consumption Consumption) Exclusive() Consumption {
	consumption.IsExclusive = true
	return consumption
}

func (consumption Consumption) NoLocal() Consumption {
	consumption.IsNoLocal = true
	return consumption
}

func (consumption Consumption) NoWait() Consumption {
	consumption.IsNoWait = true
	return consumption
}