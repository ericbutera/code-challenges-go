package ratelimit

import (
	"sync"
	"time"
)

// note: use a real library https://github.com/RussellLuo/slidingwindow

type SlidingWindowRateLimiter struct {
	mu           sync.Mutex
	windowSize   time.Duration
	maxRequests  int
	requestTimes []time.Time
}

// NewSlidingWindowRateLimiter initializes a new sliding window rate limiter.
func NewSlidingWindowRateLimiter(windowSize time.Duration, maxRequests int) *SlidingWindowRateLimiter {
	return &SlidingWindowRateLimiter{ //nolint:exhaustruct
		windowSize:   windowSize,
		maxRequests:  maxRequests,
		requestTimes: []time.Time{},
	}
}

// Allow checks if a request can be made within the rate limit.
func (rl *SlidingWindowRateLimiter) Allow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	cutoff := now.Add(-rl.windowSize)

	// Remove timestamps outside the sliding window
	i := 0
	for ; i < len(rl.requestTimes) && rl.requestTimes[i].Before(cutoff); i++ { //nolint:revive
	}
	rl.requestTimes = rl.requestTimes[i:]

	// Check if the request can be allowed
	if len(rl.requestTimes) < rl.maxRequests {
		rl.requestTimes = append(rl.requestTimes, now)
		return true
	}

	return false
}
