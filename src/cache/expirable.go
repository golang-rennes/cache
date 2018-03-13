package cache

import (
	"log"
	"time"
)

// expirable is a in-memory implementation of a cache, where content expires
// after some delay
type expirable struct {
	data map[string]*value
}

// value wraps the content, and adds an expiration timer
type value struct {
	content []byte
	timer   *time.Timer
}

const lifetime = 5 * time.Second

// NewExpirable constructs a expirable cache
func NewExpirable() T {
	return &expirable{
		data: make(map[string]*value),
	}
}

// Get retrieves a value from the cache
func (c *expirable) Get(key string) ([]byte, bool) {
	if val, ok := c.data[key]; ok {
		// extend content lifetime
		val.timer.Reset(lifetime)
		return val.content, true
	}
	return nil, false
}

// Add stores a value into the cache
func (c *expirable) Add(key string, content []byte) {
	if val, ok := c.data[key]; ok {
		// key already exists: update content and return
		val.content = content
		val.timer.Reset(lifetime)
		return
	}

	val := value{
		content: content,
		// setup expiration
		timer: time.AfterFunc(lifetime, c.expire(key)),
	}
	c.data[key] = &val
}

// Invalidate removes a single value from the cache
func (c *expirable) Invalidate(key string) {
	if val, ok := c.data[key]; ok {
		// stop expiration timer before removing value
		val.timer.Stop()
		delete(c.data, key)
	}
}

func (c *expirable) expire(key string) func() {
	return func() {
		log.Printf("Content for key %s has expired", key)
		c.Invalidate(key)
	}
}
