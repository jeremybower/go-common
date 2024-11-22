package backoff

import (
	"math/rand"
	"time"
)

// To avoid avoid overflow, 4^15 seconds = ~34 years
const maxAttempts = 15

type Backoff struct {
	minRetryDelaySecs int32
	maxRetryDelaySecs int32
	resetDelaySecs    int32
	jitter            bool
	attempt           int32
	retryAt           time.Time
	resetAt           time.Time
}

func New(
	minRetryDelaySecs int32,
	maxRetryDelaySecs int32,
	resetDelaySecs int32,
	jitter bool,
) *Backoff {
	return &Backoff{
		minRetryDelaySecs: minRetryDelaySecs,
		maxRetryDelaySecs: maxRetryDelaySecs,
		resetDelaySecs:    resetDelaySecs,
		jitter:            jitter,
		attempt:           0,
		retryAt:           time.Time{},
		resetAt:           time.Time{},
	}
}

func (b *Backoff) Reset() {
	b.attempt = 0
	b.retryAt = time.Time{}
	b.resetAt = time.Time{}
}

func (b *Backoff) Attempt(now time.Time) time.Duration {
	if now.After(b.resetAt) || now.Equal(b.resetAt) {
		b.attempt = 0
	}

	if now.Before(b.retryAt) {
		return b.retryAt.Sub(now)
	}

	delaySecs := pow(4, max(int32(0), min(b.attempt, maxAttempts)))
	if b.minRetryDelaySecs > 0 {
		delaySecs = max(delaySecs, b.minRetryDelaySecs)
	}
	if b.maxRetryDelaySecs > 0 {
		delaySecs = min(delaySecs, b.maxRetryDelaySecs)
	}

	retryAfter := time.Duration(delaySecs) * time.Second
	if b.jitter {
		retryAfter = time.Duration(float64(retryAfter) * (1.1 - (rand.Float64() * 0.2)))
	}

	resetAfter := time.Duration(b.resetDelaySecs) * time.Second

	b.attempt++
	b.retryAt = now.Add(retryAfter)
	b.resetAt = now.Add(resetAfter)
	return retryAfter
}

func (b *Backoff) Wait() <-chan time.Time {
	return b.wait(time.Now())
}

func (b *Backoff) wait(now time.Time) <-chan time.Time {
	return time.After(b.Attempt(now))
}

func pow(n, m int32) int32 {
	if m == 0 {
		return 1
	}

	if m == 1 {
		return n
	}

	result := n
	for i := int32(2); i <= m; i++ {
		result *= n
	}

	return result
}
