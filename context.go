package messenger

type Context interface {
	GetMessage() Message

	Acknowledge() error
	NegativeAcknowledge(requeue bool) error

	Publish(message Message) error
}
