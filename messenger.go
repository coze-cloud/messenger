package messenger

type Messenger interface {
	Publish(exchange Exchange, queue Queue, message Message) error
	Consume(exchange Exchange, queue Queue, consumer Consumer) error
}
