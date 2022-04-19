# messenger
📬 A minimalistic library for abstracting asynchronus messaging (inspired by [Vice](https://github.com/matryer/vice))

## Installation

Adding *messenger* to your Go module is as easy as calling this command in your project

```shell
go get github.com/coze-cloud/messenger
```

## Usage

Being a minimalistic library, *messenger* only provides the basics. In essence the basics are transport layer abstractions and a simple (but opinonated) message implementation. The later however is completly independend from the transport layer. Thus you can easily build your own way of handling messages!

Currently supported transports are:
- In-Memory / Local
- RabitMQ / AMQP
- Redis Streams

### Setting up a messenger

```go
ctx, cancel := context.WithCancel(context.Background)
defer cancel()

msgr := local.NewLocalMessenger(ctx)
```

### Receiving data

Receiving and sending data utilizes Go channels, which work just as you expect them to!     
> The abstraction and thus the use of exchange and consumer in messenger is heavily inspired by the way messages are processed by RabbitMQ using the [Publish/Subscribe](https://www.rabbitmq.com/tutorials/tutorial-three-go.html) transport.

```go
data, ok := <- msgr.Receive("exchange", "consumer")
if !ok {
    // Channel is closed, so no further data will arrive
    return
}
fmt.Println(string(data))
```

### Sending data

```go
msgr.Send("exchange") <- []byte("Hello world!")
```

>**IMPORTANT:** Do not close send channels after no furhter transmission will take place, as they are shared instances acros the life of your program.

---

Copyright © 2022 - The cozy team **& contributors**
