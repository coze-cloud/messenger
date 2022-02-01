package messenger

type localConsumer struct {
	queue *Queue
	Name  string
}

func NewLocalConsumer(queue *Queue, name string) *localConsumer {
	return &localConsumer{
		queue: queue,
		Name:  name,
	}
}

func (c *localConsumer) Consume(handler handler) {

}

func (c *localConsumer) Close() {
}
