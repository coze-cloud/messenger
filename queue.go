package messenger

import (
	"github.com/mcuadros/go-defaults"
)

type Queue struct {
	Name                 string `default:"default"`
	IsDurable            bool   `default:"false"`
	ShouldDeleteIfUnused bool   `default:"false"`
	IsExclusive          bool   `default:"false"`
	IsNoWait             bool   `default:"false"`
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

func (queue Queue) Durable() Queue {
	queue.IsDurable = true
	return queue
}

func (queue Queue) DeleteIfUnused() Queue {
	queue.ShouldDeleteIfUnused = true
	return queue
}

func (queue Queue) Exclusive() Queue {
	queue.IsExclusive = true
	return queue
}

func (queue Queue) NoWait() Queue {
	queue.IsNoWait = true
	return queue
}
