package cache

import "testing"

var (
	key     = "test"
	content = []byte("content content content")
)

type cacheTest struct {
	T
	name string
}

var caches = []*cacheTest{
	//	&cacheTest{name: "None", T: NewNone()},
	&cacheTest{name: "Memory", T: NewMemory()},
}

func TestCacheAddGet(t *testing.T) {
	for _, c := range caches {
		t.Run(c.name, func(t *testing.T) {
			c.Add(key, content)
			cachedContent, ok := c.Get(key)
			if !ok {
				t.Errorf("Content not found for key %s", key)
			} else if string(cachedContent) != string(content) {
				t.Errorf("Content does not match: got %s, expected %s", string(cachedContent), string(content))
			}
		})
	}
}

func TestCacheInvalidate(t *testing.T) {
	for _, c := range caches {
		t.Run(c.name, func(t *testing.T) {
			c.Add(key, content)
			c.Invalidate(key)
			_, ok := c.Get(key)
			if ok {
				t.Errorf("Content for key %s was not invalidated", key)
			}
		})
	}
}
