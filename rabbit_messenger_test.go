package messenger

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockConnection struct {
	mock.Mock
}

func (c *mockConnection) GetChannel() (Channel, error) {
	args := c.Called()
	return args.Get(0).(Channel), args.Error(1)
}

func (c *mockConnection) Close() error {
	args := c.Called()
	return args.Error(0)
}

func (c *mockChannel) Publish(exchange Exchange, queue Queue, message Message) error {
	args := c.Called(exchange, queue, message)
	return args.Error(0)
}

func (c *mockChannel) Consume(queue Queue, consumer Consumer) error {
	args := c.Called(queue, consumer)
	return args.Error(0)
}

func (c *mockChannel) Close() error {
	args := c.Called()
	return args.Error(0)
}

func TestRabbitMessenger_Publish(t *testing.T) {
	// Arrange
	exchange := NewExchange()

	queue := NewQueue()
	args := make(map[string]interface{})

	address := newAddress()
	message := NewMessage("Hello World")

	channel := &mockChannel{}
	channel.On("DeclareQueue", queue, args).
		Return(nil)
	channel.On("Publish", exchange, queue, message.SendFrom(address)).
		Return(nil).
		Once()
	channel.On("Close").
		Return(nil).
		Once()

	connection := &mockConnection{}
	connection.On("GetChannel").
		Return(channel, nil).
		Once()

	messenger := rabbitMessenger{
		address: address,
		connection: connection,
	}

	// Act
	err := messenger.Publish(exchange, queue, message)

	// Assert
	assert.NoError(t, err)
	channel.AssertExpectations(t)
	connection.AssertExpectations(t)
}

func TestRabbitMessenger_Consume(t *testing.T) {
	// Arrange
	exchange := NewExchange()

	queue := NewQueue()
	args := make(map[string]interface{})

	address := newAddress()
	consumer := NewConsumer(nil)

	channel := &mockChannel{}
	channel.On("DeclareQueue", queue, args).
		Return(nil)
	channel.On("Consume", queue, consumer.locatedAt(address)).
		Return( nil).
		Once()
	channel.On("Close").
		Return(nil).
		Once()

	connection := &mockConnection{}
	connection.On("GetChannel").
		Return(channel, nil).
		Once()

	messenger := rabbitMessenger{
		address: address,
		connection: connection,
	}

	// Act
	free, err := messenger.Consume(exchange, queue, consumer)

	// Assert
	assert.NoError(t, err)
	assert.NoError(t, free())
	channel.AssertExpectations(t)
	connection.AssertExpectations(t)
}