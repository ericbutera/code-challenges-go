package ratelimit_test

import (
	"testing"
	"time"

	"github.com/ericbutera/code-challenges-go/patterns/ratelimit"
	"github.com/stretchr/testify/assert"
)

func TestSlidingWindowRateLimiter(t *testing.T) {
	t.Parallel()
	t.Run("allow requests within rate limit", func(t *testing.T) {
		t.Parallel()
		limiter := ratelimit.NewSlidingWindowRateLimiter(1*time.Second, 3)

		assert.True(t, limiter.Allow(), "First request should be allowed")
		assert.True(t, limiter.Allow(), "Second request should be allowed")
		assert.True(t, limiter.Allow(), "Third request should be allowed")
		assert.False(t, limiter.Allow(), "Fourth request should be denied as it exceeds the limit")
	})

	t.Run("allow requests after window passes", func(t *testing.T) {
		t.Parallel()
		limiter := ratelimit.NewSlidingWindowRateLimiter(1*time.Second, 3)

		assert.True(t, limiter.Allow(), "First request should be allowed")
		assert.True(t, limiter.Allow(), "Second request should be allowed")
		assert.True(t, limiter.Allow(), "Third request should be allowed")
		assert.False(t, limiter.Allow(), "Fourth request should be denied")

		time.Sleep(1 * time.Second)

		assert.True(t, limiter.Allow(), "Request should be allowed after window passes")
	})

	t.Run("handles high concurrency", func(t *testing.T) {
		t.Parallel()
		limiter := ratelimit.NewSlidingWindowRateLimiter(1*time.Second, 5)

		results := make(chan bool, 10)
		for i := 0; i < 10; i++ {
			go func() {
				results <- limiter.Allow()
			}()
		}

		allowed := 0
		for i := 0; i < 10; i++ {
			if <-results {
				allowed++
			}
		}

		assert.Equal(t, 5, allowed, "Only 5 requests should be allowed within the limit")
	})
}
