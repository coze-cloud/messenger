package messenger

type rabbitMessenger struct {
	address *address
	connection Connection
}

func NewRabbitMessenger(url string) (Messenger, error) {
	connection, err := newRabbitConnection(url)
	if err != nil {
		return nil, err
	}

	return &rabbitMessenger{
		address: newAddress(),
		connection: connection,
	}, nil
}

func (m rabbitMessenger) Publish(exchange Exchange, queue Queue, message Message) error {
	channel, err := m.prepare(exchange, queue)
	defer func(channel Channel) {
		_ = channel.Close()
	}(channel)
	if err != nil {
		return err
	}

	if err := channel.Publish(exchange, queue, message.SendFrom(m.address)); err != nil {
		return err
	}

	if queue.autoRemove {
		if err := channel.DeleteQueue(queue); err != nil {
			return err
		}
	}

	return nil
}

func (m rabbitMessenger) Consume(exchange Exchange, queue Queue, consumer Consumer) (func() error, error) {
	channel, err := m.prepare(exchange, queue)
	if err != nil {
		return channel.Close, err
	}

	if err := channel.Consume(exchange, queue, consumer.locatedAt(m.address)); err != nil {
		return channel.Close, err
	}

	return channel.Close, nil
}

func (m rabbitMessenger) prepare(exchange Exchange, queue Queue) (Channel, error) {
	channel, err := m.connection.GetChannel()
	if err != nil {
		return channel, err
	}

	var commands []Command

	commands = append(commands, newCreateRabbitExchangeCommand(channel, exchange))
	commands = append(commands, newCreateRabbitQueueCommand(channel, queue))
	commands = append(commands, newBindRabbitQueueExchangeCommand(channel, exchange, queue))

	for _, command := range commands {
		if err := command.Handle(); err != nil {
			return nil, err
		}
	}

	return channel, nil
}

func (m rabbitMessenger) Close(handler func (err error)) {
	if err := m.connection.Close(); handler != nil && err != nil {
		handler(err)
	}
}