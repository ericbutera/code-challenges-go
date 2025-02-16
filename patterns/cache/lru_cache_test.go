// https://leetcode.com/problems/lru-cache/solutions/1581028/go-using-standard-library
// https://github.com/hashicorp/golang-lru
package cache_test

import (
	"container/list"
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

type LRUCache struct {
	cache    map[string]*list.Element
	linklist *list.List
	capacity int
	mu       sync.Mutex // Optimize with RWMutex. allows multiple readers or a single writer
}

type cacheEntry struct {
	key   string
	value any
}

func NewLRU(capacity int) LRUCache {
	return LRUCache{
		cache:    make(map[string]*list.Element, capacity),
		linklist: list.New(),
		capacity: capacity,
	}
}

func (c *LRUCache) Get(key string) any {
	// check if exists
	// if so move to front and return value
	// else return nil
	c.mu.Lock()
	defer c.mu.Unlock()

	// key exists, update most recently used
	if elem, ok := c.cache[key]; ok {
		c.linklist.MoveToFront(elem)
		if entry, ok := c.getCacheEntry(elem); ok {
			return entry.value
		} // delete c.cache[key] if not a cacheEntry
	}

	return nil
}

func (c *LRUCache) Put(key string, value any) {
	// if exists, update value and move to front, exit
	// if at capacity, remove last element
	// add new element to front
	c.mu.Lock()
	defer c.mu.Unlock()

	// key is already cached, remove existing entry and add new entry to front
	if elem, ok := c.cache[key]; ok {
		if entry, valid := c.getCacheEntry(elem); valid {
			entry.value = value
			c.linklist.MoveToFront(elem)
			return
		}
	}

	// capacity reached, evict last element (least recently used) to make space
	if len(c.cache) >= c.capacity {
		c.removeOldest()
	}

	// add value as most recently used
	c.cache[key] = c.linklist.PushFront(&cacheEntry{
		key,
		value,
	})
}

func (c *LRUCache) removeOldest() {
	// must be called with lock held!
	last := c.linklist.Back()
	if last != nil {
		if entry, valid := c.getCacheEntry(last); valid {
			delete(c.cache, entry.key)
		}
		c.linklist.Remove(last)
	}
}

func (c *LRUCache) getCacheEntry(elem *list.Element) (*cacheEntry, bool) {
	if entry, ok := elem.Value.(*cacheEntry); ok {
		return entry, true
	}
	return nil, false
}

func TestLRU(t *testing.T) {
	t.Parallel()

	t.Run("Get non-existent entry", func(t *testing.T) {
		t.Parallel()
		c := NewLRU(2)
		assert.Nil(t, c.Get("1"))
	})

	t.Run("Get entry", func(t *testing.T) {
		t.Parallel()
		c := NewLRU(2)
		c.Put("1", "value1")
		assert.Equal(t, "value1", c.Get("1"))
	})

	t.Run("Get updates most recently used", func(t *testing.T) {
		t.Parallel()
		c := NewLRU(2)
		c.Put("1", "value1")
		c.Put("2", 2)
		c.Get("1")
		c.Put("3", "value3")
		assert.Nil(t, c.Get("2"))
	})

	t.Run("Put entries", func(t *testing.T) {
		t.Parallel()
		c := NewLRU(2)
		c.Put("1", "value1")
		c.Put("2", 2)
		assert.Equal(t, "value1", c.Get("1"))
		assert.Equal(t, 2, c.Get("2"))
	})

	t.Run("Eviction", func(t *testing.T) {
		t.Parallel()
		c := NewLRU(2)
		c.Put("1", "one")
		c.Put("2", "two")
		c.Put("3", "three")
		// should have 2, 3
		assert.Equal(t, "three", c.Get("3"))
		assert.Equal(t, "two", c.Get("2"))
		assert.Nil(t, c.Get("1")) // 1 should be evicted
	})

	t.Run("Update entry sets most recently used", func(t *testing.T) {
		t.Parallel()
		c := NewLRU(2)
		c.Put("1", "one")
		c.Put("2", "two")
		c.Put("1", "newone") // update 1 making its access time greater than 2
		c.Put("3", "three")  // evict 2
		assert.Equal(t, "newone", c.Get("1"))
		assert.Equal(t, "three", c.Get("3"))
		assert.Nil(t, c.Get("2"))
	})

	t.Run("Update entry updates value", func(t *testing.T) {
		t.Parallel()
		c := NewLRU(2)
		c.Put("1", "one")
		c.Put("1", "newone")
		assert.Equal(t, "newone", c.Get("1"))
	})

	t.Run("Concurrent access", func(t *testing.T) {
		// go test -race ./lru_cache_test.go
		t.Parallel()

		c := NewLRU(5)
		var wg sync.WaitGroup
		numGoroutines := 100

		// Perform concurrent writes
		wg.Add(numGoroutines)
		for i := 0; i < numGoroutines; i++ {
			go func(i int) {
				defer wg.Done()
				c.Put(fmt.Sprintf("key-%d", i%5), i) // Cycle through 5 keys
			}(i)
		}

		// Perform concurrent reads
		wg.Add(numGoroutines)
		for i := 0; i < numGoroutines; i++ {
			go func(i int) {
				defer wg.Done()
				_ = c.Get(fmt.Sprintf("key-%d", i%5))
			}(i)
		}

		wg.Wait()

		// Final validation: Ensure we can retrieve values for existing keys
		for i := 0; i < 5; i++ {
			val := c.Get(fmt.Sprintf("key-%d", i))
			assert.NotNil(t, val)
		}
	})
}

// Unsafe implementation missing locks
type LRUCacheUnsafe struct {
	cache    map[string]*list.Element
	linklist *list.List
	capacity int
}

func NewLRUUnsafe(capacity int) LRUCacheUnsafe {
	return LRUCacheUnsafe{
		cache:    make(map[string]*list.Element, capacity),
		linklist: list.New(),
		capacity: capacity,
	}
}

func (c *LRUCacheUnsafe) Get(key string) any {
	// key exists, update most recently used
	if elem, ok := c.cache[key]; ok {
		c.linklist.MoveToFront(elem)
		if entry, ok := elem.Value.(*cacheEntry); ok {
			return entry.value
		} // delete c.cache[key] if not a cacheEntry
	}
	return nil
}

func (c *LRUCacheUnsafe) Put(key string, value any) {
	if elem, ok := c.cache[key]; ok {
		// key is already cached, remove existing entry and add new entry to front
		c.linklist.Remove(elem)
	} else if len(c.cache) >= c.capacity {
		// capacity reached, evict last element (least recently used) to make space
		last := c.linklist.Back()
		if last != nil {
			val := c.linklist.Remove(last)
			if evicted, ok := val.(*cacheEntry); ok {
				delete(c.cache, evicted.key)
			}
		}
	}

	// add new element to front
	c.cache[key] = c.linklist.PushFront(&cacheEntry{
		key,
		value,
	})
}
