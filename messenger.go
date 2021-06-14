package messenger

type Messenger interface {
	GetAddress() address

	Publish(queue Queue, publication Publication) error
	Consume(queue Queue, consumption Consumption) error
}