package messenger

import (
	"log"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	service, _ := NewRabbitMessenger("amqp://guest:guest@localhost:5672/")

	defaultExchange := NewExchange()
	helloQueue := NewQueue().Named("hello").WithTimeToLive(10 * time.Second)

	helloMessage := NewMessage("Hello World")
	helloPublication := NewPublication(helloMessage)

	_ = service.Publish(defaultExchange, helloQueue, helloPublication)

	helloConsumption := NewConsumption(func(ctx Context) {
		log.Println(ctx.GetMessage())
	}).Named("helloConsumer").AutoAcknowledge()

	_ = service.Consume(defaultExchange, helloQueue, helloConsumption)

	forever := make(chan bool); <-forever
}
