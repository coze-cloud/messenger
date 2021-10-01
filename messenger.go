package messenger

type Messenger interface {
	GetAddress() address

	Publish(exchange Exchange, queue Queue, publication Publication) error
	Consume(exchange Exchange, queue Queue, consumption Consumption) (func() error, error)

	Close(func (err error))
}