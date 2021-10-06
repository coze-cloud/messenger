package messenger

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func (c *mockChannel) BindQueueToExchange(exchange Exchange, queue Queue) error {
	args := c.Called(exchange, queue)
	return args.Error(0)
}

func TestBindRabbitQueueExchangeCommand_Handle(t *testing.T) {
	// Arrange
	exchange := NewExchange()

	queue := NewQueue()

	channel := &mockChannel{}

	command := newBindRabbitQueueExchangeCommand(channel, exchange, queue)

	// Act
	err := command.Handle()

	// Assert
	assert.NoError(t, err)
	channel.AssertNotCalled(t, "BindQueueToExchange")
}

func TestBindRabbitQueueExchangeCommand_Handle_WithNamedExchange(t *testing.T) {
	// Arrange
	exchange := NewExchange().Named("ExampleExchange")

	queue := NewQueue()

	channel := &mockChannel{}
	channel.On("BindQueueToExchange", exchange, queue).
		Return(nil).
		Once()

	command := newBindRabbitQueueExchangeCommand(channel, exchange, queue)

	// Act
	err := command.Handle()

	// Assert
	assert.NoError(t, err)
	channel.AssertExpectations(t)
}
