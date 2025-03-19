package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	Entries  map[string]cacheEntry
	Mu       *sync.RWMutex
	Interval time.Duration
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		Entries:  make(map[string]cacheEntry),
		Mu:       &sync.RWMutex{},
		Interval: interval,
	}

	cache.reapLoop()
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	c.Entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.Mu.RLock()
	defer c.Mu.RUnlock()
	entry, ok := c.Entries[key]
	if !ok {
		return nil, false
	}
	if time.Since(entry.createdAt) > c.Interval {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reap() {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	for key, entry := range c.Entries {
		if time.Since(entry.createdAt) > c.Interval {
			delete(c.Entries, key)
		}
	}
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.Interval)
	go func() {
		for range ticker.C {
			c.reap()
		}
	}()
}
