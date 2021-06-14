package messenger

import (
	"encoding/json"
	uuid "github.com/satori/go.uuid"
	"github.com/streadway/amqp"
	"time"
)

type rabbitHandlerContext struct {
	Context

	receiver address
	channel  *amqp.Channel
	delivery amqp.Delivery
}

func newRabbitHandlerContext(receiver address, channel *amqp.Channel, delivery amqp.Delivery) *rabbitHandlerContext {
	context := new(rabbitHandlerContext)
	context.receiver = receiver
	context.channel = channel
	context.delivery = delivery

	return context
}

func (context rabbitHandlerContext) GetMessage() Message {
	message := Message{}
	_ =  json.Unmarshal(context.delivery.Body, &message)

	series := uuid.UUID{}
	_ = json.Unmarshal([]byte(context.delivery.Headers["Series"].(string)), &series)
	message.Series = series
	message.Revision = context.delivery.Headers["Revision"].(int)
	from := address{}
	_ = json.Unmarshal([]byte(context.delivery.Headers["From"].(string)), &from)
	message.From = from
	timeStamp := time.Time{}
	_ = json.Unmarshal([]byte(context.delivery.Headers["TimeStamp"].(string)), &timeStamp)
	message.TimeStamp = timeStamp

	return message.ReceivedFrom(context.receiver)
}

func (context rabbitHandlerContext) Acknowledge() error {
	return context.delivery.Ack(false)
}

func (context rabbitHandlerContext) NegativeAcknowledge(requeue bool) error {
	return context.delivery.Nack(false, requeue)
}

func (context rabbitHandlerContext) Publish(queue Queue, publication Publication) error {
	rabbitQueue, err := NewRabbitQueueFactory(context.channel, queue).Produce()
	if err != nil {
		return err
	}
	return newRabbitPublisher(context.receiver, context.channel, publication).Publish(rabbitQueue)
}