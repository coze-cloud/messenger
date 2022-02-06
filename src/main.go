package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/cozy-hosting/messenger/src/local"
)

func main() {
	messenger := local.NewLocalMessenger()
	go func() {
		panic(<-messenger.Errors())
	}()

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			message := <-messenger.Receive("test", "foo1")
			fmt.Println("1: " + string(message))
			if string(message) == "quit" {
				return
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			message := <-messenger.Receive("test", "foo2")
			fmt.Println("2: " + string(message))
			if string(message) == "quit" {
				return
			}
		}
	}()

	go func() {
		defer wg.Done()
		for {
			message := <-messenger.Receive("test", "foo2")
			fmt.Println("3: " + string(message))
			if string(message) == "quit" {
				return
			}
		}
	}()

	for i := 0; i < 10; i++ {
		messenger.Send("test") <- []byte(fmt.Sprintf("%d", i))
		time.Sleep(500 * time.Millisecond)
	}
	messenger.Send("test") <- []byte("quit")
	wg.Wait()
}
