package messenger

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"time"
)

func TestNewMessage(t *testing.T) {
	// Arrange
	body := "Hello World"
	bodyType := reflect.TypeOf(body).Name()

	// Act
	message := NewMessage(body)

	// Assert
	assert.NotNil(t, message.series)
	assert.Equal(t, message.revision, 0)

	assert.Nil(t, message.from)
	assert.Nil(t, message.to)

	assert.True(t, time.Now().UTC().After(message.timeStamp))

	assert.Equal(t, message.body, body)
	assert.Equal(t, message.bodyType, bodyType)
}

func TestMessage_ReplyTo(t *testing.T) {
	// Arrange
	message := NewMessage("Hello World")

	replyBody := "Hello ReplyTo"
	replyBodyType := reflect.TypeOf(replyBody).Name()

	// Act
	replyMessage := message.ReplyTo(replyBody)

	// Assert
	assert.Equal(t, replyMessage.series, replyMessage.series)
	assert.Equal(t, replyMessage.revision, message.revision + 1)

	assert.True(t, replyMessage.timeStamp.After(message.timeStamp))

	assert.Equal(t, replyMessage.body, replyBody)
	assert.Equal(t, replyMessage.bodyType, replyBodyType)
}

func TestMessage_SendFrom(t *testing.T) {
	// Arrange
	message := NewMessage("Hello World")
	from := &address{}

	// Act
	result := message.SendFrom(from)

	// Assert
	assert.Equal(t, result.from, from)
}

func TestMessage_ReceivedBy(t *testing.T) {
	// Arrange
	message := NewMessage("Hello World")
	to := &address{}

	// Act
	result := message.ReceivedBy(to)

	// Assert
	assert.Equal(t, result.to, to)
}