package messenger

import (
	"log"
	"sync"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	service, _ := NewRabbitMessenger("amqp://guest:guest@localhost:5672/")
	defer service.Close(func(err error) {
		log.Fatal(err)
	})

	defaultExchange := NewExchange().Named("test")
	helloQueue := NewQueue().Named("hello").WithTimeToLive(10 * time.Second)

	helloMessage := NewMessage("Hello World")
	helloPublication := NewPublication(helloMessage)

	_ = service.Publish(defaultExchange, helloQueue, helloPublication)

	block := sync.WaitGroup{}
	block.Add(2)

	helloReplyQueue := NewQueue().WithTopic("reply")
	helloReplyConsumption := NewConsumption(func(ctx Context) {
		log.Println(ctx.GetMessage())
		block.Done()
	})

	_ = service.Consume(defaultExchange, helloReplyQueue, helloReplyConsumption)

	helloConsumption := NewConsumption(func(ctx Context) {
		log.Println(ctx.GetMessage())
		block.Done()

		reply := ctx.GetMessage().Reply("Hello as well")
		_ = ctx.Publish(defaultExchange, NewQueue().WithTopic("reply"), NewPublication(reply))
	}).Named("helloConsumer").AutoAcknowledge()

	_ = service.Consume(defaultExchange, helloQueue, helloConsumption)

	block.Wait()
}
