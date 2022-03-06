package messenger

type Messenger interface {
	Receive(exchange string, name string) <-chan []byte
	Send(exchange string) chan<- []byte

	Errors() <-chan error

	Stop()
}
