package messenger

type Connection interface {
	GetChannel() (Channel, error)

	Close() error
}
