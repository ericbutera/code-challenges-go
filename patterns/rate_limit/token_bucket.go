package rate_limit

import (
	"sync"
	"time"
)

// note: use a real library https://github.com/uber-go/ratelimit

type TokenBucket struct {
	capacity   int           // Maximum tokens in the bucket
	tokens     int           // Current tokens in the bucket
	fillRate   time.Duration // Time to generate one token
	lastRefill time.Time     // Last time the bucket was refilled
	mu         sync.Mutex
}

// NewTokenBucket initializes a new token bucket.
func NewTokenBucket(capacity int, fillRate time.Duration) *TokenBucket {
	return &TokenBucket{
		capacity:   capacity,
		tokens:     capacity,
		fillRate:   fillRate,
		lastRefill: time.Now(),
	}
}

// Allow attempts to consume one token from the bucket. Returns true if successful, false otherwise.
func (tb *TokenBucket) Allow() bool {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	tb.refill()

	if tb.tokens > 0 {
		tb.tokens--
		return true
	}

	return false
}

// refill adds tokens to the bucket based on the elapsed time since the last refill.
func (tb *TokenBucket) refill() {
	now := time.Now()
	elapsed := now.Sub(tb.lastRefill)
	tokensToAdd := int(elapsed / tb.fillRate)

	if tokensToAdd > 0 {
		tb.tokens = min(tb.capacity, tb.tokens+tokensToAdd)
		tb.lastRefill = now
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
