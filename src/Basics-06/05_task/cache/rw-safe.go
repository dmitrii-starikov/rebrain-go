package cache

import "sync"

type SafeCache struct {
	storage map[string]int
	mu      sync.RWMutex
}

func (c *SafeCache) Increase(key string, value int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.storage[key] += value
}

func (c *SafeCache) Set(key string, value int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.storage[key] = value
}

func (c *SafeCache) Get(key string) int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.storage[key]
}

func (c *SafeCache) Remove(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.storage, key)
}

func NewSafeCache() *SafeCache {
	return &SafeCache{
		storage: make(map[string]int),
	}
}
