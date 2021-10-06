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
	assert.False(t, uuid.Equal(uuid.UUID{}, address.Id))
	assert.Equal(t, name, address.Name)
}

func TestAddress_String(t *testing.T) {
	// Arrange
	address := newAddress()

	expected := fmt.Sprintf("%s(%s)",
		strings.Split(address.Id.String(), "-")[0],
		address.Name,
	)

	// Act
	result := address.String()

	// Assert
	assert.Equal(t, expected, result)
}