package messenger

type Command interface {
	Handle() error
}
