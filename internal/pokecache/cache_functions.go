package pokecache

import (
	"sync"
	"time"
)

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		entries: map[string]CacheEntry{},
		mu:      &sync.Mutex{},
	}

	// Note to future self:
	// It is likely this will cause you problems
	go cache.reapLoop(interval)
	return cache
}

func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[key] = CacheEntry{
		createdAt: time.Now().UTC(),
		val:       value,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	value, exists := c.entries[key]
	return value.val, exists
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for range ticker.C {
		c.mu.Lock()
		for k, v := range c.entries {
			if time.Now().UTC().Sub(v.createdAt) > interval {
				delete(c.entries, k)
			}
		}
		c.mu.Unlock()
	}
}
