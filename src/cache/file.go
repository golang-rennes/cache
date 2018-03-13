package cache

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// Filenamer allows to get a file name from an object
type Filenamer interface {
	GetFilename(key string) string
}

// FilenamerCache combines the Filenamer and cache interfaces by aggregating them
type FilenamerCache interface {
	T
	Filenamer
}

// file is a filesystem implementation of a cache, which also satisfies the
// Filenamer interface
type file struct {
	dirname string
}

// NewFile constructs a file cache
func NewFile(dirname string) T {
	return &file{
		dirname: dirname,
	}
}

// Get retrieves a value from the cache
func (c *file) Get(key string) ([]byte, bool) {
	content, err := ioutil.ReadFile(filepath.Join(c.dirname, key))
	if err != nil {
		return nil, false
	}
	return content, true
}

// Add stores a value into the cache
func (c *file) Add(key string, content []byte) {
	ioutil.WriteFile(filepath.Join(c.dirname, key), content, 0644)
}

// Invalidate removes a single value from the cache
func (c *file) Invalidate(key string) {
	os.Remove(filepath.Join(c.dirname, key))
}

// GetFilename returns the file name containing the content
func (c *file) GetFilename(key string) string {
	return filepath.Join(c.dirname, key)
}
