package messenger

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestNewConsumer(t *testing.T) {
	// Arrange
	handler := func(ctx Context) {}

	// Act
	consumer := NewConsumer(handler)

	// Assert
	assert.Equal(t, reflect.ValueOf(handler).Pointer(), reflect.ValueOf(consumer.handler).Pointer())
	assert.Equal(t, "", consumer.name)
	assert.False(t, consumer.autoAcknowledge)
}

func TestConsumer_Name(t *testing.T) {
	// Arrange
	name := "ExampleConsumer"

	// Act
	consumer := NewConsumer(nil).Named(name)

	// Assert
	assert.Equal(t, name, consumer.name)
}

func TestConsumer_ShouldAutoAcknowledges(t *testing.T) {
	// Arrange

	// Act
	consumer := NewConsumer(nil).ShouldAutoAcknowledge()

	// Assert
	assert.True(t, consumer.autoAcknowledge)
}