package backoff

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAttempt(t *testing.T) {
	t.Parallel()

	b := New(0, 0, 120, false)

	now := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	assert.Equal(t, 1*time.Second, b.Attempt(now))

	now = now.Add(1 * time.Second)
	assert.Equal(t, 4*time.Second, b.Attempt(now))

	now = now.Add(4 * time.Second)
	assert.Equal(t, 16*time.Second, b.Attempt(now))

	now = now.Add(16 * time.Second)
	assert.Equal(t, 64*time.Second, b.Attempt(now))

	now = now.Add(60 * time.Second) // before retry delay
	assert.Equal(t, 4*time.Second, b.Attempt(now))

	now = now.Add(120 * time.Second) // after reset delay
	assert.Equal(t, 1*time.Second, b.Attempt(now))
}

func TestReset(t *testing.T) {
	t.Parallel()

	b := New(0, 0, 120, false)

	now := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	assert.Equal(t, 1*time.Second, b.Attempt(now))

	now = now.Add(1 * time.Second)
	assert.Equal(t, 4*time.Second, b.Attempt(now))

	b.Reset()
	assert.Equal(t, 1*time.Second, b.Attempt(now))
}

func TestAttemptConstrained(t *testing.T) {
	t.Parallel()

	b := New(10, 30, 60, false)

	now := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	assert.Equal(t, 10*time.Second, b.Attempt(now))

	now = now.Add(10 * time.Second)
	assert.Equal(t, 10*time.Second, b.Attempt(now))

	now = now.Add(10 * time.Second)
	assert.Equal(t, 16*time.Second, b.Attempt(now))

	now = now.Add(16 * time.Second)
	assert.Equal(t, 30*time.Second, b.Attempt(now))
}

func TestAttemptJitter(t *testing.T) {
	t.Parallel()

	b := New(0, 0, 60, true)

	now := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	delay := b.Attempt(now)
	assert.InDelta(t, 1*time.Second, delay, float64(time.Second)*0.1)

	now = now.Add(delay)
	delay = b.Attempt(now)
	assert.InDelta(t, 4*time.Second, delay, float64(4*time.Second)*0.1)

	now = now.Add(delay)
	delay = b.Attempt(now)
	assert.InDelta(t, 16*time.Second, delay, float64(16*time.Second)*0.1)

	now = now.Add(delay)
	delay = b.Attempt(now)
	assert.InDelta(t, 64*time.Second, delay, float64(64*time.Second)*0.1)
}

func TestWait(t *testing.T) {
	t.Parallel()

	b := New(0, 0, 60, false)
	ch := b.Wait()
	assert.NotNil(t, ch)
}
