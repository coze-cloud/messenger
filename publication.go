package messenger

import "github.com/mcuadros/go-defaults"

type publication struct {
	Message     Message
	Exchange    string `default:""`
	IsMandatory bool   `default:"false"`
	IsImmediate bool   `default:"false"`
}

func NewPublication(message Message) publication {
	publication := new(publication)
	defaults.SetDefaults(publication)
	publication.Message = message

	return *publication
}

func (publication publication) AtExchange(exchange string) publication {
	publication.Exchange = exchange
	return publication
}

func (publication publication) Mandatory() publication {
	publication.IsMandatory = true
	return publication
}

func (publication publication) Immediate() publication {
	publication.IsImmediate = true
	return publication
}