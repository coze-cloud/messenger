package messenger

type Consumer interface {
	Consume(handler handler)
}
