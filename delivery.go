package messenger

type Delivery interface {
	GetMessage() (Message, error)

	Acknowledge() error
	NegativeAcknowledge(requeue bool) error
}
