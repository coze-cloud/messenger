package messenger

import (
	uuid "github.com/satori/go.uuid"
	"reflect"
	"time"
)

type Message struct {
	series uuid.UUID
	revision int

	from *address
	to *address

	timeStamp time.Time

	bodyType string
	body interface{}
}

// BEGIN: Constructor

func NewMessage(body interface{}) Message {
	message := new(Message)
	message.series = uuid.NewV4()

	message.timeStamp = time.Now().UTC()

	message.body = body
	message.bodyType = reflect.TypeOf(body).Name()

	return *message
}

// END: Constructor

// BEGIN: Methods

func (message Message) ReplyTo(body interface{}) Message {
	message.revision++

	message.timeStamp = time.Now().UTC()

	message.body = body
	message.bodyType = reflect.TypeOf(body).Name()

	return message
}

func (message Message) SendFrom(from *address) Message {
	message.from = from

	return message
}

func (message Message) ReceivedBy(to *address) Message {
	message.to = to

	return message
}

// END: Methods