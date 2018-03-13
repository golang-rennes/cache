package cache

import "testing"

var (
	key     = "test"
	content = []byte("content content content")
)

func TestCacheAddGet(t *testing.T) {
	c := NewNone()

	c.Add(key, content)
	cachedContent, ok := c.Get(key)
	if ok && string(cachedContent) != string(content) {
		t.Errorf("Content does not match: got %s, expected %s", string(cachedContent), string(content))
	}
}

func TestCacheInvalidate(t *testing.T) {
	c := NewNone()

	c.Add(key, content)
	c.Invalidate(key)
	_, ok := c.Get(key)
	if ok {
		t.Errorf("Content for key %s was not invalidated", key)
	}
}
