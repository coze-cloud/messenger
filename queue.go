package messenger

import (
	"github.com/mcuadros/go-defaults"
)

type queue struct {
	Name                 string `default:"default"`
	IsDurable            bool   `default:"false"`
	ShouldDeleteIfUnused bool   `default:"false"`
	IsExclusive          bool   `default:"false"`
	IsNoWait             bool   `default:"false"`
}

func NewQueue() queue {
	queue := new(queue)
	defaults.SetDefaults(queue)

	return *queue
}

func (queue queue) Named(name string) queue {
	queue.Name = name
	return queue
}

func (queue queue) Durable() queue {
	queue.IsDurable = true
	return queue
}

func (queue queue) DeleteIfUnused() queue {
	queue.ShouldDeleteIfUnused = true
	return queue
}

func (queue queue) Exclusive() queue {
	queue.IsExclusive = true
	return queue
}

func (queue queue) NoWait() queue {
	queue.IsNoWait = true
	return queue
}
