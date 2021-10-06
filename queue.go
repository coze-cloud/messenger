package messenger

import "time"

type Queue struct {
	name string
	topic string
	autoRemove bool
	timeToLive uint64
}

func NewQueue() Queue {
	return Queue{autoRemove: true}
}

func (q Queue) Named(name string) Queue {
	q.name = name
	q.topic = name
	q.autoRemove = false

	return q
}

func (q Queue) WithTopic(topic string) Queue {
	q.topic = topic

	return q
}

func (q Queue) ShouldAutoRemove() Queue {
	q.autoRemove = true

	return q
}

func (q Queue) WithTimeToLive(duration time.Duration) Queue {
	q.timeToLive = uint64(duration.Milliseconds())

	return q
}