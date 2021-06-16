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
	_ =  json.Unmarshal(context.delivery.Body, &message.Body)

	series := uuid.UUID{}
	_ = json.Unmarshal(context.delivery.Headers["Series"].([]uint8), &series)
	message.Series = series
	message.Revision = int(context.delivery.Headers["Revision"].(int32))
	from := address{}
	_ = json.Unmarshal(context.delivery.Headers["From"].([]uint8), &from)
	message.From = from
	timeStamp := time.Time{}
	_ = json.Unmarshal(context.delivery.Headers["TimeStamp"].([]uint8), &timeStamp)
	message.TimeStamp = timeStamp

	return message.ReceivedFrom(context.receiver)
}

func (context rabbitHandlerContext) Acknowledge() error {
	return context.delivery.Ack(false)
}

func (context rabbitHandlerContext) NegativeAcknowledge(requeue bool) error {
	return context.delivery.Nack(false, requeue)
}

func (context rabbitHandlerContext) Publish(exchange Exchange, queue Queue, publication Publication) error {
	return newRabbitPublisher(context.receiver, context.channel, publication).Publish(exchange, queue)
}