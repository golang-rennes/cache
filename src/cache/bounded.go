package cache

import (
	"container/heap"
	"log"
	"time"
)

// bounded is a bounded-size in-memory implementation of a cache
type bounded struct {
	data           map[string]*item
	insertionQueue *itemHeap
}

type item struct {
	key            string
	content        []byte
	index          int
	lastAccessTime time.Time
}

const maxItems = 3

// itemHeap implements a priority queue, using the built-in heap package
type itemHeap []*item

func (h itemHeap) Len() int { return len(h) }
func (h itemHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].index = j
	h[j].index = i
}

// Less defines the removal policy, which is always a trade-off between
// use-case, efficiency, simplicity.
// Currently, remove the least recently used item.
// See https://en.wikipedia.org/wiki/Cache_replacement_policies for alternatives.
func (h itemHeap) Less(i, j int) bool { return h[i].lastAccessTime.Sub(h[j].lastAccessTime) < 0 }

func (h *itemHeap) Push(x interface{}) {
	n := len(*h)
	i := x.(*item)
	i.lastAccessTime = time.Now()
	i.index = n
	*h = append(*h, i)
}

func (h *itemHeap) Pop() interface{} {
	old := *h
	n := len(old)
	i := old[n-1]
	i.index = -1
	*h = old[0 : n-1]
	return i
}

// NewBounded constructs a bounded cache
func NewBounded() T {
	insertionQueue := make(itemHeap, 0, maxItems)
	heap.Init(&insertionQueue)

	return &bounded{
		data:           make(map[string]*item, maxItems),
		insertionQueue: &insertionQueue,
	}
}

// Get retrieves a item from the cache
func (c *bounded) Get(key string) ([]byte, bool) {
	if i, ok := c.data[key]; ok {
		// update last access date and return content
		i.lastAccessTime = time.Now()
		heap.Fix(c.insertionQueue, i.index)
		return i.content, true
	}
	return nil, false
}

// Add stores a item into the cache
func (c *bounded) Add(key string, content []byte) {
	if i, ok := c.data[key]; ok {
		// key already exists: update content and return
		i.content = content
		i.lastAccessTime = time.Now()
		heap.Fix(c.insertionQueue, i.index)
		return
	}

	if len(c.data) >= maxItems {
		// remove least recently used element
		x := heap.Pop(c.insertionQueue)
		i := x.(*item)
		log.Printf("Content for key %s has been replaced", i.key)
		delete(c.data, i.key)
	}

	// add new element
	i := item{
		key:     key,
		content: content,
	}
	heap.Push(c.insertionQueue, &i)
	c.data[key] = &i
}

// Invalidate removes a single item from the cache
func (c *bounded) Invalidate(key string) {
	if i, ok := c.data[key]; ok {
		heap.Remove(c.insertionQueue, i.index)
		delete(c.data, key)
	}
}
