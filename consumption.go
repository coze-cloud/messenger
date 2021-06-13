package messenger

import "github.com/mcuadros/go-defaults"

type consumption struct {
	Handler           handler
	Name              string `default:""`
	IsAutoAcknowledge bool `default:"false"`
	IsExclusive       bool `default:"false"`
	IsNoLocal         bool `default:"false"`
	IsNoWait          bool `default:"false"`
}

func NewConsumption(handler handler) consumption {
	consumption := new(consumption)
	defaults.SetDefaults(consumption)
	consumption.Handler = handler

	return *consumption
}

func (consumption consumption) Named(name string) consumption {
	consumption.Name = name
	return consumption
}

func (consumption consumption) AutoAcknowledge() consumption {
	consumption.IsAutoAcknowledge = true
	return consumption
}

func (consumption consumption) Exclusive() consumption {
	consumption.IsExclusive = true
	return consumption
}

func (consumption consumption) NoLocal() consumption {
	consumption.IsNoLocal = true
	return consumption
}

func (consumption consumption) NoWait() consumption {
	consumption.IsNoWait = true
	return consumption
}