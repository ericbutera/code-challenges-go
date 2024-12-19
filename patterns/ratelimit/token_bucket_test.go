package ratelimit_test

import (
	"testing"
	"time"

	"github.com/ericbutera/code-challenges-go/patterns/ratelimit"
	"github.com/stretchr/testify/assert"
)

func TestTokenBucket(t *testing.T) {
	t.Parallel()
	t.Run("allow tokens within capacity", func(t *testing.T) {
		t.Parallel()
		limiter := ratelimit.NewTokenBucket(5, time.Second)

		for i := 0; i < 5; i++ {
			assert.True(t, limiter.Allow(), "Should allow tokens within capacity")
		}

		assert.False(t, limiter.Allow(), "Should not allow more tokens than capacity")
	})

	t.Run("refill tokens over time", func(t *testing.T) {
		t.Parallel()
		limiter := ratelimit.NewTokenBucket(3, 500*time.Millisecond)

		for i := 0; i < 3; i++ {
			assert.True(t, limiter.Allow(), "Should allow tokens within capacity")
		}

		assert.False(t, limiter.Allow(), "Should not allow tokens after capacity is depleted")

		time.Sleep(1 * time.Second)
		assert.True(t, limiter.Allow(), "Should allow a token after refill")
		assert.True(t, limiter.Allow(), "Should allow another token after refill")
		assert.False(t, limiter.Allow(), "Should not allow more tokens than refill rate")
	})

	t.Run("refill respects capacity limit", func(t *testing.T) {
		t.Parallel()
		limiter := ratelimit.NewTokenBucket(3, 200*time.Millisecond)

		for i := 0; i < 3; i++ {
			assert.True(t, limiter.Allow(), "Should allow tokens within capacity")
		}

		time.Sleep(1 * time.Second)
		for i := 0; i < 3; i++ {
			assert.True(t, limiter.Allow(), "Should allow refilled tokens up to capacity")
		}

		assert.False(t, limiter.Allow(), "Should not allow more tokens than capacity even after a long wait")
	})
}
