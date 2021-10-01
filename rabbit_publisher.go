package messenger

import (
	"encoding/json"
	"github.com/streadway/amqp"
)

type rabbitPublisher struct {
	sender      address
	channel     *amqp.Channel
	publication Publication
}

func newRabbitPublisher(sender address, channel *amqp.Channel, publication Publication) rabbitPublisher {
	publisher := new(rabbitPublisher)
	publisher.sender = sender
	publisher.channel = channel
	publisher.publication = publication

	return *publisher
}

func (publisher rabbitPublisher) Publish(exchange Exchange, queue Queue) error {
	message := publisher.publication.Message.SentFrom(publisher.sender)

	series, err := json.Marshal(message.Series)
	if err != nil {
		return err
	}
	from, err := json.Marshal(message.From)
	if err != nil {
		return err
	}
	timeStamp, err := json.Marshal(message.TimeStamp)
	if err != nil {
		return err
	}

	body, err := json.Marshal(message.Body)
	if err != nil {
		return err
	}

	return publisher.channel.Publish(
		exchange.Name,
		queue.Topic,
		false,
		false,
		amqp.Publishing{
			Headers: map[string]interface{}{
				"Id": message.Id,
				"Series": series,
				"Revision": message.Revision,
				"From": from,
				"TimeStamp": timeStamp,
			},
			ContentType: "text/json",
			Body:        body,
		},
	)
}