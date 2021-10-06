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
	assert.NotNil(t, message.Series)
	assert.Equal(t, message.Revision, 0)

	assert.Nil(t, message.From)
	assert.Nil(t, message.To)

	assert.True(t, time.Now().UTC().After(message.TimeStamp))

	assert.Equal(t, message.Body, body)
	assert.Equal(t, message.BodyType, bodyType)
}

func TestMessage_ReplyTo(t *testing.T) {
	// Arrange
	message := NewMessage("Hello World")

	replyBody := "Hello ReplyTo"
	replyBodyType := reflect.TypeOf(replyBody).Name()

	// Act
	time.Sleep(1 * time.Millisecond)
	replyMessage := message.ReplyTo(replyBody)

	// Assert
	assert.Equal(t, replyMessage.Series, replyMessage.Series)
	assert.Equal(t, replyMessage.Revision, message.Revision + 1)

	assert.True(t, replyMessage.TimeStamp.After(message.TimeStamp))

	assert.Equal(t, replyMessage.Body, replyBody)
	assert.Equal(t, replyMessage.BodyType, replyBodyType)
}

func TestMessage_SendFrom(t *testing.T) {
	// Arrange
	message := NewMessage("Hello World")
	from := &address{}

	// Act
	result := message.SendFrom(from)

	// Assert
	assert.Equal(t, result.From, from)
}

func TestMessage_ReceivedBy(t *testing.T) {
	// Arrange
	message := NewMessage("Hello World")
	to := &address{}

	// Act
	result := message.ReceivedBy(to)

	// Assert
	assert.Equal(t, result.To, to)
}