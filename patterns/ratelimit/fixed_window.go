package ratelimit

import (
	"sync"
	"time"
)

// note: use a real library https://github.com/mennanov/limiters

type FixedWindowLimiter struct {
	windowSize time.Duration
	maxHits    int
	windows    sync.Map // maps window start time to count of hits
}

// NewFixedWindowLimiter initializes a new FixedWindowLimiter.
func NewFixedWindowLimiter(windowSize time.Duration, maxHits int) *FixedWindowLimiter {
	return &FixedWindowLimiter{
		windowSize: windowSize,
		maxHits:    maxHits,
	}
}

// Allow checks if a request is allowed.
func (l *FixedWindowLimiter) Allow() bool {
	currentWindow := time.Now().Unix() / int64(l.windowSize.Seconds())

	// Clean up old windows to save memory
	l.windows.Range(func(key, _ interface{}) bool {
		if k, ok := key.(int64); ok && k < currentWindow {
			l.windows.Delete(key)
		}
		return true
	})

	// Check and update the current window
	value, _ := l.windows.LoadOrStore(currentWindow, 0)
	currentCount, ok := value.(int)
	if !ok {
		currentCount = 0
	}
	if currentCount < l.maxHits {
		l.windows.Store(currentWindow, currentCount+1)
		return true
	}

	return false
}
