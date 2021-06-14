package messenger

type Context interface {
	GetMessage() Message

	Acknowledge() error
	NegativeAcknowledge(requeue bool) error
	Publish(queue Queue, publication Publication) error
}