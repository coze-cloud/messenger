package messenger

type Context interface {
	GetDelivery() Delivery

	Publish(message Message) error
}
