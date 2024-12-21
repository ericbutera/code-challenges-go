package ratelimit_test

import (
	"testing"
	"time"

	"github.com/ericbutera/code-challenges-go/patterns/ratelimit"
	"github.com/stretchr/testify/assert"
)

func TestFixedWindowLimiter(t *testing.T) {
	t.Parallel()
	limiter := ratelimit.NewFixedWindowLimiter(1*time.Second, 5)

	for i := 0; i < 5; i++ {
		assert.True(t, limiter.Allow())
	}

	assert.False(t, limiter.Allow())

	time.Sleep(1 * time.Second) // wait for new window

	assert.True(t, limiter.Allow())
}

func TestDeny(t *testing.T) {
	t.Parallel()
	limiter := ratelimit.NewFixedWindowLimiter(1*time.Second, 0)
	assert.False(t, limiter.Allow())
}

func TestDeny1(t *testing.T) {
	t.Parallel()
	limiter := ratelimit.NewFixedWindowLimiter(1*time.Second, 1)
	assert.True(t, limiter.Allow())
	assert.False(t, limiter.Allow())
}
