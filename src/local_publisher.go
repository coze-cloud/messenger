package messenger

type localPublisher struct {
	exchange *Exchange
}

func NewLocalPublisher(exchange *Exchange) *localPublisher {
	return &localPublisher{
		exchange: exchange,
	}
}

func (p *localPublisher) Publish(message Message) {

}
