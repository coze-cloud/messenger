package messenger

type Channel interface {
	DeclareExchange(exchange Exchange) error

	DeclareQueue(queue Queue, args map[string]interface{}) error
	DeleteQueue(queue Queue) error

	BindQueueToExchange(exchange Exchange, queue Queue) error

	Publish(exchange Exchange, queue Queue, message Message) error
	Consume(queue Queue, consumer Consumer) error
}