package postgres

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTime(t *testing.T) {
	t.Parallel()

	original := time.Date(2024, 1, 2, 3, 4, 5, int(time.Millisecond+time.Microsecond+time.Nanosecond), time.UTC)
	expected := time.Date(2024, 1, 2, 3, 4, 5, int(time.Millisecond+time.Microsecond), time.UTC)
	assert.Equal(t, expected, Time(original))
}
