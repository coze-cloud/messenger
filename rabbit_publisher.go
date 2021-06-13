package messenger

import (
	"encoding/json"
	"github.com/streadway/amqp"
)

type rabbitPublisher struct {
	sender      address
	channel     *amqp.Channel
	publication publication
}

func newRabbitPublisher(sender address, channel *amqp.Channel, publication publication) rabbitPublisher {
	publisher := new(rabbitPublisher)
	publisher.sender = sender
	publisher.channel = channel
	publisher.publication = publication

	return *publisher
}

func (publisher rabbitPublisher) Publish(queue amqp.Queue) error {
	message := publisher.publication.Message.SendFrom(publisher.sender)
	body, err := json.Marshal(message.Body)
	if err != nil {
		return err
	}

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

	return publisher.channel.Publish(
		publisher.publication.Exchange,
		queue.Name,
		publisher.publication.IsMandatory,
		publisher.publication.IsImmediate,
		amqp.Publishing{
			Headers: map[string]interface{}{
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