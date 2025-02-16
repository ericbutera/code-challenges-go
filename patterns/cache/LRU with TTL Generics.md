# LRU, TTL, Generics

This allows it to store any key-value pair type while maintaining the same functionality, including TTL support.

## Key Changes

- Generic Parameters `[K comparable, V any]`:
  - `K` (key) must be `comparable` (to support `map` lookup).
  - `V` (value) can be of any type.
- Updated Methods to Use Generics: The cache now works for any key-value type, making it reusable.

## Generic LRU Cache Implementation

```go
package main

import (
  "container/list"
  "sync"
  "time"
)

type LRUCache[K comparable, V any] struct {
  cache    map[K]*list.Element
  linklist *list.List
  capacity int
  mu       sync.Mutex
  ttl      time.Duration
}

type cacheEntry[K comparable, V any] struct {
  key        K
  value      V
  expiration time.Time
}

func NewLRU[K comparable, V any](capacity int, ttl time.Duration) *LRUCache[K, V] {
  return &LRUCache[K, V]{
    cache:    make(map[K]*list.Element, capacity),
    linklist: list.New(),
    capacity: capacity,
    ttl:      ttl,
  }
}

func (c *LRUCache[K, V]) Get(key K) (V, bool) {
  c.mu.Lock()
  defer c.mu.Unlock()

  var zeroValue V // Default zero value for type V

  if elem, ok := c.cache[key]; ok {
    entry, valid := c.getCacheEntry(elem)
    if valid {
      if time.Now().After(entry.expiration) {
        // Entry has expired, remove it
        c.removeElement(elem)
        return zeroValue, false
      }
      // Move to front and return value
      c.linklist.MoveToFront(elem)
      return entry.value, true
    }
  }
  return zeroValue, false
}

func (c *LRUCache[K, V]) Put(key K, value V) {
  c.mu.Lock()
  defer c.mu.Unlock()

  // If key exists, update value and move to front
  if elem, ok := c.cache[key]; ok {
    if entry, valid := c.getCacheEntry(elem); valid {
      entry.value = value
      entry.expiration = time.Now().Add(c.ttl) // Refresh TTL
      c.linklist.MoveToFront(elem)
      return
    }
  }

  // Evict expired entries before adding a new one
  c.evictExpired()

  // If capacity exceeded, evict least recently used
  if len(c.cache) >= c.capacity {
    c.removeOldest()
  }

  // Insert new entry
  entry := &cacheEntry[K, V]{
    key:        key,
    value:      value,
    expiration: time.Now().Add(c.ttl),
  }
  c.cache[key] = c.linklist.PushFront(entry)
}

func (c *LRUCache[K, V]) removeOldest() {
  // must be called with lock held!
  last := c.linklist.Back()
  if last != nil {
    c.removeElement(last)
  }
}

func (c *LRUCache[K, V]) evictExpired() {
  for e := c.linklist.Back(); e != nil; {
    prev := e.Prev()
    if entry, valid := c.getCacheEntry(e); valid {
      if time.Now().After(entry.expiration) {
        c.removeElement(e)
      }
    }
    e = prev
  }
}

func (c *LRUCache[K, V]) removeElement(elem *list.Element) {
  if entry, valid := c.getCacheEntry(elem); valid {
    delete(c.cache, entry.key)
  }
  c.linklist.Remove(elem)
}

func (c *LRUCache[K, V]) getCacheEntry(elem *list.Element) (*cacheEntry[K, V], bool) {
  if entry, ok := elem.Value.(*cacheEntry[K, V]); ok {
    return entry, true
  }
  return nil, false
}
```

---

### Usage Example

```go
package main

import (
  "fmt"
  "time"
)

func main() {
  cache := NewLRU[string, int](3, 5*time.Second) // String keys, integer values

  cache.Put("a", 1)
  cache.Put("b", 2)
  cache.Put("c", 3)

  val, ok := cache.Get("a") // Should return 1, true
  fmt.Println(val, ok)

  time.Sleep(6 * time.Second) // Wait for TTL expiration

  val, ok = cache.Get("a") // Should return 0, false (expired)
  fmt.Println(val, ok)

  cache.Put("d", 4) // Should evict least recently used key

  val, ok = cache.Get("b") // Likely evicted if LRU applied correctly
  fmt.Println(val, ok)
}
```

### Benefits of Using Generics

- Supports any key type (`string`, `int`, etc.): `K comparable` ensures keys can be used in maps.
- Supports any value type (`int`, `string`, structs, etc.): `V any` makes it flexible.
- Zero-value return handling: When `Get()` fails, it returns the zero value of `V` (e.g., `0` for `int`, `""` for `string`, `nil` for pointers).

This version is type-safe, efficient, and maintains all the LRU + TTL behavior. Let me know if you need modifications!
