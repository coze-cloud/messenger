package messenger

type Queue struct {
	Name     string
	messages chan *Message
}

func NewQueue(name string) *Queue {
	queue := &Queue{
		Name:     name,
		messages: make(chan *Message),
	}
	return queue
}

func (q *Queue) SendMessage(message *Message) {
	q.messages <- message
}

func (q *Queue) ReceiveMessage() *Message {
	return <-q.messages
}

func (q *Queue) Close() {
	close(q.messages)
}
