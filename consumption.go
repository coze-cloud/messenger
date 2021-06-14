package messenger

import "github.com/mcuadros/go-defaults"

type Consumption struct {
	Handler           handler
	Name              string `default:""`
	IsAutoAcknowledge bool `default:"false"`
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