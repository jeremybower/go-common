package env

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPort(t *testing.T) {
	assert.NoError(t, IsPort("80"))
	assert.ErrorIs(t, IsPort("0"), ErrInvalidPortNumber)
	assert.ErrorIs(t, IsPort("65536"), ErrInvalidPortNumber)
	assert.ErrorIs(t, IsPort("invalid"), ErrInvalidPortNumber)
}

func TestNotEmpty(t *testing.T) {
	assert.NoError(t, NotEmpty("test"))
	assert.ErrorIs(t, NotEmpty(""), ErrEmpty)
}

func TestNonNegative(t *testing.T) {
	assert.NoError(t, NonNegative(0))
	assert.NoError(t, NonNegative(1))
	assert.ErrorIs(t, NonNegative(-1), ErrNonNegative)
}

func TestPositive(t *testing.T) {
	assert.NoError(t, Positive(1))
	assert.ErrorIs(t, Positive(0), ErrPositive)
	assert.ErrorIs(t, Positive(-1), ErrPositive)
}
