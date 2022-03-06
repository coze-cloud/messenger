package messenger

import "context"

func Consume(ctx context.Context, receiver <-chan []byte, consumer func(message Message)) {
	go func() {
		for {
			select {
			case data := <-receiver:
				message, err := Deserialize(data)
				if err != nil {
					continue
				}
				consumer(message)
			case <-ctx.Done():
				return
			}
		}
	}()
}
