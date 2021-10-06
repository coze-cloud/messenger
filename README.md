# messenger
📬 A minimalistic library for abstracting asynchronus messaging

## Installation

Adding *messenger* to your Go module is as easy as calling this command in your project

```shell
go get github.com/cozy-hosting/messenger
```

## Usage

Being a minimalistic library, *messenger* only provides the basics. The rest is up to your specific need.

### Create a messenger

```go
msgr, err := messenger.NewRabbitMessenger("amqp://guest:guest@localhost:5672/")
defer service.Close(func(err error) {
    log.Fatal(err)
})
if err != nil {
    log.Fatal(err)
}
```

### Define an exchange

```go
defaultExchange := messenger.NewExchange()
```

### Define a queue

```go
exampleQueue := messenger.NewQueue().Named("example")
```

### Publish a message

```go
helloMessage := messenger.NewMessage("Hello World")

if err := msgr.Publish(defaultExchange, exampleQueue, helloMessage); err != nil {
    log.Fatal(err)
}
```

### Consume messages

```go
helloConsumption := messenger.NewConsumption(func(ctx m.Context) {
    message, err := ctx.GetDelivery().GetMessage()
    if err != err {
        log.Fatal(err)
    }   
    
    log.Println(message)
}).AutoAcknowledge()

free, err := msgr.Consume(defaultExchange, exampleQueue, helloConsumption)
defer free()
if err != nil {
    log.Fatal(err)
}
```

## Future plans

* [x] Unit tests for the existing components
* [ ] Support for more message queue implementations

---

Copyright © 2021 - The cozy team **& contributors**
