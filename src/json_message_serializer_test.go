package messenger

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJsonMessageSerializer_Serialize(t *testing.T) {
	// Arrange
	message := NewMessage("Hello World")
	serializer := newJsonMessageSerializer(message)

	serializedMessage, _ := json.Marshal(message)

	// Act
	result, err := serializer.Serialize()

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, string(serializedMessage), result)
}

func TestJsonMessageDeserializer_Deserialize(t *testing.T) {
	// Arrange
	message := NewMessage("Hello World")
	serializedMessage, _ := json.Marshal(message)
	deserializer := newJsonMessageDeserializer(string(serializedMessage))

	// Act
	result, err := deserializer.Deserialize()

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, message, result)
}