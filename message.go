package messenger

import (
	"encoding/json"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"reflect"
	"strings"
	"time"
)

type Message struct {
	Series   uuid.UUID `json:"series"`
	Revision int       `json:"revision"`

	From *address `json:"from"`
	To   *address `json:"to"`

	TimeStamp time.Time `json:"time_stamp"`

	BodyType string      `json:"body_type"`
	Body     interface{} `json:"body"`
}

// BEGIN: Constructor

func NewMessage(body interface{}) Message {
	return Message{
		Series:    uuid.NewV4(),
		TimeStamp: time.Now().UTC(),
		BodyType:  reflect.TypeOf(body).String(),
		Body:      body,
	}
}

// END: Constructor

// BEGIN: Methods

func (m Message) ReplyTo(body interface{}) Message {
	m.Revision++

	m.TimeStamp = time.Now().UTC()

	m.Body = body
	m.BodyType = reflect.TypeOf(body).Name()

	return m
}

func (m Message) SendFrom(from *address) Message {
	m.From = from
	return m
}

func (m Message) ReceivedBy(to *address) Message {
	m.To = to

	return m
}

func (message Message) String() string {
	timeStamp := message.TimeStamp.Format(time.RFC3339)
	body, _ := json.Marshal(message.Body)

	return fmt.Sprintf("%s.%d(%s) @ %s, %s -> %s, %s",
		strings.Split(message.Series.String(), "-")[0],
		message.Revision,
		message.BodyType,
		timeStamp,
		message.From,
		message.To,
		body,
	)
}

// END: Methods
