package messenger

import (
	"errors"
	"reflect"

	"github.com/mitchellh/mapstructure"
)

var (
	UnexpectedMessageType = errors.New("unexpected message type")
	BodyCanNotBeDecoded   = errors.New("unable to decode message body")
)

type Handler[T any] func(handleable T) error

func Handle[T any](message Message, handler Handler[T]) error {
	var handleable T
	if message.Type != reflect.TypeOf(handleable).String() {
		return UnexpectedMessageType
	}
	if err := mapstructure.Decode(message.Body, &handleable); err != nil {
		return BodyCanNotBeDecoded
	}
	return handler(handleable)
}
