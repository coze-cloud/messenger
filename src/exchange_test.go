package messenger_test

import (
	"context"
	"testing"

	messenger "github.com/cozy-hosting/messenger/src"
	"github.com/stretchr/testify/assert"
)

func TestExchangeBindQueueDoesBindSuccessfully(t *testing.T) {
	// Arrange
	exchange := messenger.NewExchange("test_exchange")
	queue := messenger.NewQueue("test_queue", 10)

	// Act
	exchange.BindQueue(queue)

	// Assert
	assert.Equal(t, 1, len(exchange.Queues))
}

func TestExchangeUnbindQueueDoesUnbinSuccessfully(t *testing.T) {
	// Arrange
	exchange := messenger.NewExchange("test_exchange")
	queue := messenger.NewQueue("test_queue", 10)
	exchange.BindQueue(queue)

	// Act
	exchange.UnbindQueue(queue)

	// Assert
	assert.Equal(t, 0, len(exchange.Queues))
}

func TestExchangeSendMessageIsReceivedByAnyBoundQueue(t *testing.T) {
	// Arrange
	exchange := messenger.NewExchange("test_exchange")
	queue1 := messenger.NewQueue("test_queue_1", 10)
	queue2 := messenger.NewQueue("test_queue_2", 10)
	exchange.BindQueue(queue1)
	exchange.BindQueue(queue2)

	message := messenger.NewMessage("Hello word")

	// Act
	exchange.SendMessage(&message)

	// Assert
	receivedMessage1, err := queue1.ReceiveMessage(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, message, *receivedMessage1)

	receivedMessage2, err := queue2.ReceiveMessage(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, message, *receivedMessage2)
}
