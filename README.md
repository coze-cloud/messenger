# messenger
📬 A minimalistic library for abstracting asynchronus messaging

## Installation

Adding *messenger* to your Go module is as easy as calling this command in your project

```shell
go get github.com/cozy-hosting/messenger
```

## Usage

Being a minimalistic library, *messenger* only provides the basiscs. The rest is up to your specific need.

### Create a messenger

```go
service, err := messenger.NewRabbitMessenger("amqp://guest:guest@localhost:5672/")
if err != nil {
    log.Fatal(err)
}
```

### Define a queue

```go
exampleQueue := messenger.NewQueue().Named("example")
```

### Publish a message

```go
helloMessage := messenger.NewMessage("Hello World")
helloPublication := messenger.NewPublication(helloMessage)

if err := service.Publish(exampleQueue, helloPublication); err != nil {
    log.Fatal(err)
}
```

### Consume messages

```go
helloConsumption := messenger.NewConsumption(func(ctx m.Context) {
    log.Println(ctx.GetMessage().Body)
}).AutoAcknowledge()

if err := service.Consume(exampleQueue, helloConsumption); err != nil {
    log.Fatal(err)
}
```

## Future plans

* [ ] Unit tests for the existing components
* [ ] Support for more message queue implementations

---

Copyright © 2021 - The cozy team **& contributors**
