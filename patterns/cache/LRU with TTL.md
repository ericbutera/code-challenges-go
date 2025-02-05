# LRU with TTL

Here’s an updated version of your LRU cache that includes a TTL (Time-To-Live) mechanism to mimic Redis’s `allkeys-lru` and `volatile-lru` behaviors. This implementation tracks expiration times and removes expired entries upon access or insertion.

```go
import (
    "container/list"
    "sync"
    "time"
)

type LRUCache struct {
    cache    map[string]*list.Element
    linklist *list.List
    capacity int
    mu       sync.Mutex
    ttl      time.Duration
}

type cacheEntry struct {
    key        string
    value      any
    expiration time.Time
}

func NewLRU(capacity int, ttl time.Duration) LRUCache {
    return LRUCache{
        cache:    make(map[string]*list.Element, capacity),
        linklist: list.New(),
        capacity: capacity,
        ttl:      ttl,
    }
}

func (c *LRUCache) Get(key string) any {
    c.mu.Lock()
    defer c.mu.Unlock()

    if elem, ok := c.cache[key]; ok {
        entry, valid := c.getCacheEntry(elem)
        if valid {
            if time.Now().After(entry.expiration) {
                // Entry has expired, remove it
                c.removeElement(elem)
                return nil
            }
            // Move to front and return value
            c.linklist.MoveToFront(elem)
            return entry.value
        }
    }
    return nil
}

func (c *LRUCache) Put(key string, value any) {
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
    entry := &cacheEntry{
        key:        key,
        value:      value,
        expiration: time.Now().Add(c.ttl),
    }
    c.cache[key] = c.linklist.PushFront(entry)
}

func (c *LRUCache) removeOldest() {
    // must be called with lock held!
    last := c.linklist.Back()
    if last != nil {
        c.removeElement(last)
    }
}

func (c *LRUCache) evictExpired() {
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

func (c *LRUCache) removeElement(elem *list.Element) {
    if entry, valid := c.getCacheEntry(elem); valid {
        delete(c.cache, entry.key)
    }
    c.linklist.Remove(elem)
}

func (c *LRUCache) getCacheEntry(elem *list.Element) (*cacheEntry, bool) {
    if entry, ok := elem.Value.(*cacheEntry); ok {
        return entry, true
    }
    return nil, false
}
```

## Key Changes

1. Added TTL support: Each cache entry has an `expiration` field.
2. Expired Entry Handling:
   - On `Get()`, expired entries are checked and removed.
   - On `Put()`, expired entries are proactively removed.
3. Eviction Strategy:
   - The cache removes expired keys first.
   - If still over capacity, it removes the least recently used key.

This implementation mimics Redis’s `allkeys-lru` and `volatile-lru` behaviors:

- Allkeys-LRU: If the cache is full, it evicts the least recently used entry.
- Volatile-LRU: Entries with TTLs expire naturally and are evicted if accessed after expiration.

Would you like to add periodic background eviction for better memory management?
