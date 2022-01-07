package messenger

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewExchange(t *testing.T) {
	// Arrange

	// Act
	exchange := NewExchange()

	// Assert
	assert.Equal(t, "", exchange.name)
	assert.Equal(t, "direct", exchange.strategy)
	assert.False(t, exchange.autoRemove)
}

func TestExchange_Named(t *testing.T) {
	// Arrange
	name := "ExampleExchange"

	// Act
	exchange := NewExchange().Named(name)

	// Assert
	assert.Equal(t, name, exchange.name)
}

func TestExchange_WithStrategy(t *testing.T) {
	// Arrange
	strategy := "fanout"

	// Act
	exchange := NewExchange().WithStrategy(strategy)

	// Assert
	assert.Equal(t, strategy, exchange.strategy)
}

func TestExchange_ShouldAutoRemove(t *testing.T) {
	// Arrange

	// Act
	exchange := NewExchange().ShouldAutoRemove()

	// Assert
	assert.True(t, exchange.autoRemove)
}