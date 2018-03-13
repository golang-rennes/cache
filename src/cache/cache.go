// Package cache defines a cache mechanism, and several implementations of it
package cache

// T defines a cache contract, as a key-value store.
// Keys are string identifiers, values are arbitrary binary content []byte
type T interface {
	Get(key string) ([]byte, bool)
	Add(key string, content []byte)
	Invalidate(key string)
}

const cacheDir = "/tmp/cache"

// New initializes a default cache
var New = NewBounded
