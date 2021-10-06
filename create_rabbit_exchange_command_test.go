package messenger

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockChannel struct {
	mock.Mock
	Channel
}

func (c *mockChannel) DeclareExchange(exchange Exchange) error {
	args := c.Called(exchange)
	return args.Error(0)
}

func TestCreateRabbitExchangeCommand_Handle_DefaultExchange(t *testing.T) {
	// Arrange
	exchange := NewExchange()

	channel := &mockChannel{}

	command := newCreateRabbitExchangeCommand(channel, exchange)

	// Act
	err := command.Handle()

	// Assert
	assert.NoError(t, err)
	channel.AssertNotCalled(t, "DeclareExchange")
}

func TestCreateRabbitExchangeCommand_Handle_NamedExchange(t *testing.T) {
	// Arrange
	exchange := NewExchange().Named("ExampleExchange")

	channel := &mockChannel{}
	channel.On("DeclareExchange", exchange).
		Return(nil).
		Once()

	command := newCreateRabbitExchangeCommand(channel, exchange)

	// Act
	err := command.Handle()

	// Assert
	assert.NoError(t, err)
	channel.AssertExpectations(t)
}
