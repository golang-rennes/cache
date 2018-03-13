package cache

// none is a dummy implementation of a cache
type none struct{}

// NewNone constructs a none cache
func NewNone() T {
	return &none{}
}

// Get retrieves a value from the cache
func (c *none) Get(key string) ([]byte, bool) {
	return nil, false
}

// Add stores a value into the cache
func (c *none) Add(key string, content []byte) {}

// Invalidate removes a single value from the cache
func (c *none) Invalidate(key string) {}
