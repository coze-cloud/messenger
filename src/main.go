package main

import (
	"fmt"
	"time"

	"github.com/cozy-hosting/messenger/src/redis"
)

func main() {
	messenger := redis.NewRedisClient(&redis.Options{
		Network:    "tcp",
		Addr:       "host.docker.internal:6379",
		Password:   "",
		DB:         0,
		MaxRetries: 0,
	})
	defer messenger.Stop()
	go func() {
		for {
			fmt.Println(<-messenger.Errors())
		}
	}()

	go func() {
		for {
			message := <-messenger.Receive("test-*")
			fmt.Println("1: " + string(message))
		}
	}()

	go func() {
		for {
			message := <-messenger.Receive("test-*")
			fmt.Println("2: " + string(message))
		}
	}()

	for i := 0; i < 10; i++ {
		messenger.Send("test-foo") <- []byte(fmt.Sprintf("%d", i))
		time.Sleep(1000 * time.Millisecond)
	}
}
