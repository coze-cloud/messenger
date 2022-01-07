package messenger

import "github.com/streadway/amqp"

type rabbitDelivery struct {
	receiver *address
	delivery *amqp.Delivery
}

func newRabbitDelivery(receiver *address, delivery *amqp.Delivery) *rabbitDelivery {
	return &rabbitDelivery{receiver: receiver, delivery: delivery}
}

func (r rabbitDelivery) GetMessage() (Message, error) {
	deserializer := newJsonMessageDeserializer(string(r.delivery.Body))

	message, err := deserializer.Deserialize()
	if err != nil {
		return Message{}, err
	}

	return message.ReceivedBy(r.receiver), nil
}

func (r rabbitDelivery) Acknowledge() error {
	return r.delivery.Ack(false)
}

func (r rabbitDelivery) NegativeAcknowledge(requeue bool) error {
	return r.delivery.Nack(false, requeue)
}

