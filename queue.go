package messenger

import (
	"github.com/mcuadros/go-defaults"
	"time"
)

type Queue struct {
	Name             string
	Topic            string
	ShouldAutoRemove bool	`defaults:"true"`
	TimeToLive       int64
}

func NewQueue() Queue {
	queue := new(Queue)
	defaults.SetDefaults(queue)

	return *queue
}

func (queue Queue) Named(name string) Queue {
	queue.Name = name
	queue.Topic = name
	queue.ShouldAutoRemove = false
	return queue
}

func (queue Queue) WithTopic(topic string) Queue {
	queue.Topic = topic
	return queue
}

func (queue Queue) AutoRemove() Queue {
	queue.ShouldAutoRemove = true
	return queue
}

func (queue Queue) WithTimeToLive(duration time.Duration) Queue {
	queue.TimeToLive = duration.Milliseconds()
	return queue
}
