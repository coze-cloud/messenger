package consumer

import (
	"context"

	messenger "github.com/cozy-hosting/messenger/src"
)

type localConsumer struct {
	queue *messenger.Queue
}

func NewLocalConsumer(queue *messenger.Queue) *localConsumer {
	return &localConsumer{
		queue: queue,
	}
}

func (c *localConsumer) Consume(ctx context.Context, handler func(msg messenger.Message)) error {
	for {
		msg, err := c.queue.ReceiveMessage(ctx)
		if err != nil {
			return err
		}
		handler(*msg)
	}
}
