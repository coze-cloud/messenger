package messenger

import (
	"fmt"
	"reflect"
	"time"

	uuid "github.com/satori/go.uuid"
)

type Message struct {
	Id        string    `json:"id" bson:"_id"`
	Series    string    `json:"series" bson:"series"`
	Revision  int       `json:"revision" bson:"revision"`
	TimeStamp time.Time `json:"time_stamp" bson:"time_stamp"`
	Type      string    `json:"type" bson:"type"`
	Body      any       `json:"body" bson:"body"`
}

func NewMessage(body any) Message {
	series := uuid.NewV4().String()
	return Message{
		Id:        fmt.Sprintf("%s.%d", series, 0),
		Series:    series,
		Revision:  0,
		TimeStamp: time.Now().UTC(),
		Type:      reflect.TypeOf(body).String(),
		Body:      body,
	}
}

func (m Message) Reply(body any) Message {
	revision := m.Revision + 1
	return Message{
		Id:        fmt.Sprintf("%s.%d", m.Series, revision),
		Series:    m.Series,
		Revision:  revision,
		TimeStamp: time.Now().UTC(),
		Type:      reflect.TypeOf(body).String(),
		Body:      body,
	}
}
