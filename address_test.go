package messenger

import (
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNewAddress(t *testing.T) {
	// Arrange
	name, _ := os.Hostname()

	// Act
	address := newAddress()

	// Assert
	assert.False(t, uuid.Equal(uuid.UUID{}, address.id))
	assert.Equal(t, name, address.name)
}