package messenger

type rabbitMessenger struct {
	address *address
}

func NewRabbitMessenger() Messenger {
	return &rabbitMessenger{address: newAddress()}
}

func (r rabbitMessenger) Publish(exchange Exchange, queue Queue, message Message) error {
	channel, err := r.prepare(exchange, queue)
	if err != nil {
		return err
	}

	if err := channel.Publish(exchange, queue, message); err != nil {
		return err
	}

	return channel.Close()
}

func (r rabbitMessenger) Consume(exchange Exchange, queue Queue, consumer Consumer) (func() error, error) {
	channel, err := r.prepare(exchange, queue)
	if err != nil {
		return channel.Close, err
	}

	if err := channel.Consume(queue, consumer); err != nil {
		return channel.Close, err
	}

	return channel.Close, nil
}

func (r rabbitMessenger) prepare(exchange Exchange, queue Queue) (Channel, error) {
	var commands []Command

	commands = append(commands, newCreateRabbitExchangeCommand(nil, exchange))
	commands = append(commands, newCreateRabbitQueueCommand(nil, queue))
	commands = append(commands, newBindRabbitQueueExchangeCommand(nil, exchange, queue))

	for _, command := range commands {
		if err := command.Handle(); err != nil {
			return nil, err
		}
	}

	return nil, nil // TODO: Return actual channel
}
