package messenger

import (
	"fmt"
	"reflect"

	"github.com/mitchellh/mapstructure"
)

type Handler[T any] func(handleable T) error

func Handle[T any](message Message, handler Handler[T]) error {
	var handleable T
	if message.Type != reflect.TypeOf(handleable).String() {
		return fmt.Errorf("unexpected message type: %s", message.Type)
	}
	if err := mapstructure.Decode(message.Body, &handleable); err != nil {
		return fmt.Errorf("unable to decode message body: %s", err)
	}
	return handler(handleable)
}
