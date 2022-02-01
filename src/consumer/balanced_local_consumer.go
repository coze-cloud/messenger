package consumer

import (
	"context"

	messenger "github.com/cozy-hosting/messenger/src"
)

type Balancer interface {
	Pick() interface{}
	Push()
	Done()
}

type balancedLocalConsumer struct {
	balancer Balancer
	item     interface{}
	consumer *localConsumer
}

func NewBalancedLocalConsumer(balancer Balancer, item interface{}, consumer *localConsumer) *balancedLocalConsumer {
	return &balancedLocalConsumer{
		balancer: balancer,
		item:     item,
		consumer: consumer,
	}
}

func (c *balancedLocalConsumer) Consume(ctx context.Context, handler func(msg messenger.Message)) error {
	item := c.balancer.Pick()
	if item != c.item {
		c.balancer.Push()
		return nil
	}
	err := c.consumer.Consume(ctx, handler)
	c.balancer.Done()
	return err
}
