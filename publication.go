package messenger

import "github.com/mcuadros/go-defaults"

type Publication struct {
	Message     Message
}

func NewPublication(message Message) Publication {
	publication := new(Publication)
	defaults.SetDefaults(publication)
	publication.Message = message

	return *publication
}