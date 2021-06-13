# messenger
📬 A minimalistic library for abstracting asynchronus messaging

## Adding

Adding *messenger* to your Go module is as easy as calling this command in your project

```shell
go get github.com/cozy-hosting/messenger
```

## Usage

Beeing a minimalistic library, *messenger* only provides the basiscs. The rest is up to your specific need.

### Creae a messenger

```go
messenger, err := m.NewRabbitMessenger("amqp://guest:guest@localhost:5672/")
if err != nil {
    log.Fatal(err)
}
```

### Define a queue

```go
exampleQueue := m.NewQueue().Named("example")
```

### Publish a message

```go
helloMessage := m.NewMessage("Hello World")
helloPublication := m.NewPublication(helloMessage)

if err := messenger.Publish(exampleQueue, helloPublication); err != nil {
    log.Fatal(err)
}
```

### Consume messages

```go
helloConsumption := m.NewConsumption(func(ctx m.Context) {
    log.Println(ctx.GetMessage().Body)
}).AutoAcknowledge()

if err := messenger.Consume(exampleQueue, helloConsumption); err != nil {
    log.Fatal(err)
}
```

Copyright © 2021 - The cozy team **& contributors**
