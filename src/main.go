package main

import (
	"fmt"
	"time"

	"github.com/cozy-hosting/messenger/src/local"
)

func main() {
	messenger := local.NewLocalMessenger()
	defer messenger.Stop()
	go func() {
		panic(<-messenger.Errors())
	}()

	go func() {
		for {
			message := <-messenger.Receive("test", "foo")
			fmt.Println("1: " + string(message))
		}
	}()

	go func() {
		for {
			message := <-messenger.Receive("test", "foo")
			fmt.Println("2: " + string(message))
		}
	}()

	go func() {
		for {
			message := <-messenger.Receive("test", "foo")
			fmt.Println("3: " + string(message))
		}
	}()

	for i := 0; i < 100; i++ {
		messenger.Send("test") <- []byte(fmt.Sprintf("%d", i))
		time.Sleep(50 * time.Millisecond)
	}
}
