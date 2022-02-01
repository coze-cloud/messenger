package messenger

type Publisher interface {
	Publish(message Message)
}
