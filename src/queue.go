package messenger

import "context"

type Queue struct {
	Name     string
	messages chan *Message
}

func NewQueue(name string, size int) *Queue {
	queue := &Queue{
		Name:     name,
		messages: make(chan *Message, size),
	}
	return queue
}

func (q *Queue) SendMessage(message *Message) {
	q.messages <- message
}

func (q *Queue) ReceiveMessage(ctx context.Context) (*Message, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case message := <-q.messages:
		return message, nil
	}
}
