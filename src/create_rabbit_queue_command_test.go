package messenger

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func (c *mockChannel) DeclareQueue(queue Queue, args map[string]interface{}) error {
	callArgs := c.Called(queue, args)
	return callArgs.Error(0)
}

func TestCreateRabbitQueueCommand_Handle(t *testing.T) {
	// Arrange
	queue := NewQueue()

	args := make(map[string]interface{})

	channel := &mockChannel{}
	channel.On("DeclareQueue", queue, args).
		Return(nil).
		Once()

	command := newCreateRabbitQueueCommand(channel, queue)

	// Act
	err := command.Handle()

	// Assert
	assert.NoError(t, err)
	channel.AssertExpectations(t)
}

func TestCreateRabbitQueueCommand_Handle_QueueWithTTL(t *testing.T) {
	// Arrange
	ttl := 10 * time.Second
	queue := NewQueue().WithTimeToLive(ttl)

	args := make(map[string]interface{})
	args["x-message-ttl"] = queue.timeToLive

	channel := &mockChannel{}
	channel.On("DeclareQueue", queue, args).
		Return(nil).
		Once()

	command := newCreateRabbitQueueCommand(channel, queue)

	// Act
	err := command.Handle()

	// Assert
	assert.NoError(t, err)
	channel.AssertExpectations(t)
}
