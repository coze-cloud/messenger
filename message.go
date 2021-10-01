package messenger

import (
	"encoding/json"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"strconv"
	"strings"
	"time"
)

type Message struct {
	Id string `bson:"_id"`
	Series    uuid.UUID
	Revision  int
	From      address
	To        address
	TimeStamp time.Time

	Type   string
	Body   interface{}
}

func NewMessage(body interface{}) Message {
	message := new(Message)
	message.Series = uuid.NewV4()
	message.Id = message.Series.String() + "." + strconv.Itoa(message.Revision)
	message.TimeStamp = time.Now().UTC()
	message.Body = body

	return *message
}

func (message Message) OfType(messageType string) Message {
	message.Type = messageType
	return message
}

func (message Message) SentFrom(from address) Message {
	message.From = from
	return message
}

func (message Message) ReceivedFrom(from address) Message {
	message.To = from
	return message
}

func (message Message) Reply(body interface{}) Message {
	reply := NewMessage(body)
	reply.Series = message.Series
	reply.Revision = message.Revision + 1
	reply.Id = reply.Series.String() + "." + strconv.Itoa(reply.Revision)
	reply.Body = body
	return reply
}

func (message Message) String() string {
	timeStamp := message.TimeStamp.Format(time.RFC3339)
	body, _ := json.Marshal(message.Body)

	return fmt.Sprintf("%s.%d(%s) @ %s, %s -> %s, %s",
		strings.Split(message.Series.String(), "-")[0],
		message.Revision,
		message.Type,
		timeStamp,
		message.From,
		message.To,
		body,
	)
}