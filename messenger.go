package messenger

type Messenger interface {
	GetAddress() address

	Publish(queue queue, publication publication) error
	Consume(queue queue, consumption consumption) error
}