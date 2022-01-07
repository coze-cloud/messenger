package messenger

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewQueue(t *testing.T) {
	// Arrange

	// Act
	queue := NewQueue()

	// Assert
	assert.Equal(t, "", queue.name)
	assert.Equal(t, "", queue.topic)
	assert.True(t, queue.autoRemove)
	assert.Equal(t, uint64(0), queue.timeToLive)
}

func TestQueue_Named(t *testing.T) {
	// Arrange
	name := "ExampleQueue"

	// Act
	queue := NewQueue().Named(name)

	// Assert
	assert.Equal(t, name, queue.name)
}

func TestQueue_WithTopic(t *testing.T) {
	// Arrange
	topic := "ExampleTopic"

	// Act
	queue := NewQueue().WithTopic(topic)

	// Assert
	assert.Equal(t, topic, queue.topic)
}

func TestQueue_ShouldAutoRemove(t *testing.T) {
	// Arrange

	// Act
	queue := NewQueue().ShouldAutoRemove()

	// Assert
	assert.True(t, queue.autoRemove)
}

func TestQueue_WithTimeToLive(t *testing.T) {
	// Arrange
	ttl := 10 * time.Second

	// Act
	queue := NewQueue().WithTimeToLive(ttl)

	// Assert
	assert.Equal(t, uint64(ttl.Milliseconds()), queue.timeToLive)
}
