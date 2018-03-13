package cache

import "sync"

// syncMemory is a synchronous in-memory implementation of a cache
type syncMemory struct {
	data map[string][]byte
	lock sync.RWMutex
}

// NewSyncMemory constructs a syncMemory cache
func NewSyncMemory() T {
	return &syncMemory{
		data: make(map[string][]byte),
	}
}

// Get retrieves a value from the cache
func (c *syncMemory) Get(key string) ([]byte, bool) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	content, ok := c.data[key]
	return content, ok
}

// Add stores a value into the cache
func (c *syncMemory) Add(key string, content []byte) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.data[key] = content
}

// Invalidate removes a single value from the cache
func (c *syncMemory) Invalidate(key string) {
	c.lock.Lock()
	defer c.lock.Unlock()
	delete(c.data, key)
}
