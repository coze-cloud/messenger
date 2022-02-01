package messenger_test

import (
	"context"
	"testing"
	"time"

	messenger "github.com/cozy-hosting/messenger/src"
	"github.com/stretchr/testify/assert"
)

func TestQueueSendMessageMessageIsReceived(t *testing.T) {
	// Arrange
	queue := messenger.NewQueue("test", 10)
	message := messenger.NewMessage("Hello word")

	// Act
	queue.SendMessage(&message)

	// Assert
	receivedMessage, err := queue.ReceiveMessage(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, message, *receivedMessage)
}

func TestQueueReceiveMessageCanBeCanceled(t *testing.T) {
	// Arrange
	queue := messenger.NewQueue("test", 10)

	// Act
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
	defer cancel()
	_, err := queue.ReceiveMessage(ctx)

	// Assert
	assert.Error(t, err)
}
