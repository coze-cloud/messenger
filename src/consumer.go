package messenger

import "context"

type Consumer func(message Message)

func Consume(ctx context.Context, receiver <-chan []byte, consumer Consumer) {
	go func() {
		for {
			select {
			case data, ok := <-receiver:
				if !ok {
					return
				}
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
