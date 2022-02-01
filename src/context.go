package messenger

type Context struct {
	Queue   *Queue
	Message Message
}

func NewContext(queue *Queue, message Message) *Context {
	return &Context{
		Queue:   queue,
		Message: message,
	}
}
