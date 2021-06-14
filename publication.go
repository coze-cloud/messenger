package messenger

import "github.com/mcuadros/go-defaults"

type Publication struct {
	Message     Message
	Exchange    string `default:""`
	IsMandatory bool   `default:"false"`
	IsImmediate bool   `default:"false"`
}

func NewPublication(message Message) Publication {
	publication := new(Publication)
	defaults.SetDefaults(publication)
	publication.Message = message

	return *publication
}

func (publication Publication) AtExchange(exchange string) Publication {
	publication.Exchange = exchange
	return publication
}

func (publication Publication) Mandatory() Publication {
	publication.IsMandatory = true
	return publication
}

func (publication Publication) Immediate() Publication {
	publication.IsImmediate = true
	return publication
}