package publisher

import (
	messenger "github.com/cozy-hosting/messenger/src"
)

type localPublisher struct {
	exchange *messenger.Exchange
}

func NewLocalPublisher(exchange *messenger.Exchange) *localPublisher {
	return &localPublisher{}
}

func (p *localPublisher) Publish(msg messenger.Message) {
	p.exchange.SendMessage(&msg)
}
