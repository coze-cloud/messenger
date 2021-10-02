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

func NewMessage(body interface{}) Message {
	message := new(Message)
	message.series = uuid.NewV4()

	message.timeStamp = time.Now().UTC()

	message.body = body
	message.bodyType = reflect.TypeOf(body).Name()

	return *message
}

func (message Message) Reply(body interface{}) Message {
	message.revision++

	message.timeStamp = time.Now().UTC()

	message.body = body
	message.bodyType = reflect.TypeOf(body).Name()

	return message
}

func (message Message) FromSender(from *address) Message {
	message.from = from

	return message
}

func (message Message) ToReceiver(to *address) Message {
	message.to = to

	return message
}