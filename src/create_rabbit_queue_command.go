package messenger

type createRabbitQueueCommand struct {
	channel Channel
	queue Queue
}

func newCreateRabbitQueueCommand(channel Channel, queue Queue) Command {
	return &createRabbitQueueCommand{channel: channel, queue: queue}
}

func (c createRabbitQueueCommand) Handle() error {
	args := make(map[string]interface{})
	if c.queue.timeToLive > 0 {
		args["x-message-ttl"] = c.queue.timeToLive
	}

	return c.channel.DeclareQueue(c.queue, args)
}
