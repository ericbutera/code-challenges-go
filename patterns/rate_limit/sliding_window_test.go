package rate_limit_test

import (
	"testing"
	"time"

	"github.com/ericbutera/code-challenges-go/patterns/rate_limit"
	"github.com/stretchr/testify/assert"
)

func TestSlidingWindowRateLimiter(t *testing.T) {
	t.Run("allow requests within rate limit", func(t *testing.T) {
		rl := rate_limit.NewSlidingWindowRateLimiter(1*time.Second, 3)

		assert.True(t, rl.Allow(), "First request should be allowed")
		assert.True(t, rl.Allow(), "Second request should be allowed")
		assert.True(t, rl.Allow(), "Third request should be allowed")
		assert.False(t, rl.Allow(), "Fourth request should be denied as it exceeds the limit")
	})

	t.Run("allow requests after window passes", func(t *testing.T) {
		rl := rate_limit.NewSlidingWindowRateLimiter(1*time.Second, 3)

		assert.True(t, rl.Allow(), "First request should be allowed")
		assert.True(t, rl.Allow(), "Second request should be allowed")
		assert.True(t, rl.Allow(), "Third request should be allowed")
		assert.False(t, rl.Allow(), "Fourth request should be denied")

		time.Sleep(1 * time.Second)

		assert.True(t, rl.Allow(), "Request should be allowed after window passes")
	})

	t.Run("handles high concurrency", func(t *testing.T) {
		rl := rate_limit.NewSlidingWindowRateLimiter(1*time.Second, 5)

		results := make(chan bool, 10)
		for i := 0; i < 10; i++ {
			go func() {
				results <- rl.Allow()
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
