package messenger

type Balancer interface {
	Pick() interface{}
	Push()
	Done()
}
