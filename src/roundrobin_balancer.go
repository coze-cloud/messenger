package messenger

import (
	"sync"
)

type roundrobinBalancer struct {
	mutex sync.Mutex

	next  int
	items []interface{}
}

func NewRoundrobinBalancer(items []interface{}) *roundrobinBalancer {
	return &roundrobinBalancer{
		items: items,
	}
}

func (b *roundrobinBalancer) Pick() interface{} {
	if len(b.items) == 0 {
		return nil
	}

	b.mutex.Lock()

	item := b.items[b.next]
	b.next = (b.next + 1) % len(b.items)

	return item
}

func (b *roundrobinBalancer) Push() {
	b.next = (b.next + len(b.items) - 1) % len(b.items)
	b.Done()
}

func (b *roundrobinBalancer) Done() {
	b.mutex.Unlock()
}
