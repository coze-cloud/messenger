package messenger_test

import (
	"testing"
	"time"

	messenger "github.com/cozy-hosting/messenger/src"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestThatNewMessageIsValidMessage(t *testing.T) {
	// Arrange
	bodyString := "Hello world"

	// Act
	message := messenger.NewMessage(bodyString)

	// Assert
	_, err := uuid.FromString(message.Series.String())
	assert.NoError(t, err)
	assert.Equal(t, 0, message.Revision)
	assert.NotEqual(t, time.Time{}, message.TimeStamp)
	assert.Equal(t, "string", message.Type)
	assert.Equal(t, bodyString, message.Body)
}

func TestThatReplyToAMessageIsValidMessage(t *testing.T) {
	// Arrange
	message := messenger.NewMessage("Hello world")
	replyStrting := "Hello again"

	// Act
	replyMessage := message.Reply(replyStrting)

	// Assert
	assert.Equal(t, message.Series, replyMessage.Series)
	assert.Equal(t, message.Revision+1, replyMessage.Revision)
	assert.NotEqual(t, message.TimeStamp, replyMessage.TimeStamp)
	assert.Equal(t, "string", replyMessage.Type)
	assert.Equal(t, replyStrting, replyMessage.Body)
}
