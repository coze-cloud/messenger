package messenger

type Context interface {
	GetMessage() Message

	Acknowledge() error
	NegativeAcknowledge(requeue bool) error
	Publish(queue queue, publication publication) error
}