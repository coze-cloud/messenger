package messenger

import (
	"github.com/mcuadros/go-defaults"
	"time"
)

type Queue struct {
	Name             string `default:"default"`
	ShouldAutoDelete bool
	TimeToLive       int64
}

func NewQueue() Queue {
	queue := new(Queue)
	defaults.SetDefaults(queue)

	return *queue
}

func (queue Queue) Named(name string) Queue {
	queue.Name = name
	return queue
}

func (queue Queue) AutoDelete() Queue {
	queue.ShouldAutoDelete = true
	return queue
}

func (queue Queue) WithTimeToLive(duration time.Duration) Queue {
	queue.TimeToLive = duration.Milliseconds()
	return queue
}
