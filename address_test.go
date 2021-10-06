package messenger

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
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

func TestAddress_String(t *testing.T) {
	// Arrange
	address := newAddress()

	expected := fmt.Sprintf("%s(%s)",
		strings.Split(address.id.String(), "-")[0],
		address.name,
	)

	// Act
	result := address.String()

	// Assert
	assert.Equal(t, expected, result)
}