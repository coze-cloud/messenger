package messenger

type Exchange struct {
	Name   string
	Queues []*Queue
}

func NewExchange(name string) *Exchange {
	return &Exchange{
		Name:   name,
		Queues: make([]*Queue, 0),
	}
}

func (e *Exchange) BindQueue(queue *Queue) {
	e.Queues = append(e.Queues, queue)
}

func (e *Exchange) UnbindQueue(queue *Queue) {
	for i, q := range e.Queues {
		if q == queue {
			e.Queues = append(e.Queues[:i], e.Queues[i+1:]...)
			return
		}
	}
}

func (e *Exchange) SendMessage(message *Message) {
	for _, q := range e.Queues {
		q.SendMessage(message)
	}
}

func (e *Exchange) Close() {
	for _, q := range e.Queues {
		q.Close()
		e.UnbindQueue(q)
	}
}
