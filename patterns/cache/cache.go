// TTL
package cache

import (
	"sync"
	"time"
)

type MemoryCache struct {
	cache sync.Map
}

func NewInMemoryCache() *MemoryCache {
	return &MemoryCache{}
}

func (c *MemoryCache) Get(key string) (any, bool) {
	return c.cache.Load(key)
}

func (c *MemoryCache) Set(key string, value any) {
	c.cache.Store(key, value)
}

// In memory cache with expiration
type ExpireItem[T any] struct {
	value      T
	expiration time.Time
}

type ExpiringCache[T any] struct {
	cache sync.Map
	ttl   time.Duration
}

func NewExpiringCache[T any](ttl time.Duration) *ExpiringCache[T] {
	c := &ExpiringCache[T]{
		ttl: ttl,
	}
	go c.cleanupExpiredEntries()
	return c
}

func (c *ExpiringCache[T]) Get(key string) (*T, bool) {
	attempt, found := c.cache.Load(key)
	if !found {
		return nil, false
	}

	item, valid := attempt.(ExpireItem[T])
	if !valid {
		return nil, false
	}
	if time.Now().After(item.expiration) {
		c.cache.Delete(key)
		return nil, false
	}

	return &item.value, true
}

func (c *ExpiringCache[T]) Set(key string, value T) {
	expiration := time.Now().Add(c.ttl)
	c.cache.Store(key, ExpireItem[T]{
		value:      value,
		expiration: expiration,
	})
}

const CleanupIntervalFactor = 2

func (c *ExpiringCache[T]) cleanupExpiredEntries() {
	for {
		time.Sleep(c.ttl / CleanupIntervalFactor)
		now := time.Now()
		c.cache.Range(func(key, value any) bool {
			item, ok := value.(ExpireItem[T])
			if !ok {
				c.cache.Delete(key)
				return true
			}
			if now.After(item.expiration) {
				c.cache.Delete(key)
			}
			return true
		})
	}
}
