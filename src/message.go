package messenger

import (
	"reflect"
	"time"

	uuid "github.com/satori/go.uuid"
)

type Message struct {
	Series    uuid.UUID   `json:"series"`
	Revision  int         `json:"revision"`
	TimeStamp time.Time   `json:"time_stamp"`
	Type      string      `json:"type"`
	Body      interface{} `json:"body"`
}

func NewMessage(body interface{}) Message {
	return Message{
		Series:    uuid.NewV4(),
		Revision:  0,
		TimeStamp: time.Now().UTC(),
		Type:      reflect.TypeOf(body).String(),
		Body:      body,
	}
}

func (m Message) Reply(body interface{}) Message {
	m.Revision++
	m.TimeStamp = time.Now().UTC()
	m.Type = reflect.TypeOf(body).Name()
	m.Body = body

	return m
}
