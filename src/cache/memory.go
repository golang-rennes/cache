package cache

// memory is a in-memory implementation of a cache, uses a native map
type memory map[string][]byte

// NewMemory constructs a memory cache
func NewMemory() T {
	return memory(make(map[string][]byte))
}

// Get retrieves a value from the cache
func (c memory) Get(key string) ([]byte, bool) {
	content, ok := c[key]
	return content, ok
}

// Add stores a value into the cache
func (c memory) Add(key string, content []byte) {
	c[key] = content
}

// Invalidate removes a single value from the cache
func (c memory) Invalidate(key string) {
	delete(c, key)
}
