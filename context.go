package messenger

type Context interface {
	GetMessage() Message

	Acknowledge() error
	NegativeAcknowledge(requeue bool) error
	Publish(exchange Exchange, queue Queue, publication Publication) error
}