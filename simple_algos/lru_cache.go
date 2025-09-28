package simple_algos

import (
	"container/list"
	"sync"
)

type LRUCache[Key comparable, Value any] struct {
	capacity int
	cache    map[Key]*list.Element
	list     *list.List
	mu       sync.RWMutex
}

type lruCacheEntry[Key comparable, Value any] struct {
	key   Key
	value Value
}

func NewLRUCache[Key comparable, Value any](capacity int) *LRUCache[Key, Value] {
	if capacity <= 0 {
		panic("LRUCache: capacity must be > 0")
	}
	return &LRUCache[Key, Value]{
		capacity: capacity,
		cache:    make(map[Key]*list.Element),
		list:     list.New()}
}

func (lruCache *LRUCache[Key, Value]) Get(key Key) (Value, bool) {
	// if element in the list move, it to front and return it
	lruCache.mu.Lock()
	defer lruCache.mu.Unlock()

	if element, ok := lruCache.cache[key]; ok {
		lruCache.list.MoveToFront(element)
		return element.Value.(lruCacheEntry[Key, Value]).value, true
	}
	var zero Value // is nil (where nil is allowed) and 0 (false) for others like int(bool)
	return zero, false
}

func (lruCache *LRUCache[Key, Value]) Put(key Key, value Value) {
	// if element already in the list move it to front
	lruCache.mu.Lock()
	defer lruCache.mu.Unlock()
	if element, ok := lruCache.cache[key]; ok {
		lruCache.list.MoveToFront(element)
		element.Value = lruCacheEntry[Key, Value]{key, value}
		return
	}
	// Delete last element if cache is full
	if lruCache.list.Len() == lruCache.capacity {
		back := lruCache.list.Back()
		if back != nil {
			lruCache.list.Remove(back)
			delete(lruCache.cache, back.Value.(lruCacheEntry[Key, Value]).key)
		}
	}
	// push new element to front
	element := lruCache.list.PushFront(lruCacheEntry[Key, Value]{key, value})
	lruCache.cache[key] = element
}

func (lruCache *LRUCache[Key, Value]) Remove(key Key) {
	lruCache.mu.Lock()
	defer lruCache.mu.Unlock()
	if element, ok := lruCache.cache[key]; ok {
		lruCache.list.Remove(element)
		delete(lruCache.cache, key)
	}
}

func (lruCache *LRUCache[Key, Value]) Peek(key Key) (Value, bool) {
	lruCache.mu.RLock()
	defer lruCache.mu.RUnlock()
	if element, ok := lruCache.cache[key]; ok {
		return element.Value.(lruCacheEntry[Key, Value]).value, true
	}
	var zero Value
	return zero, false
}

func (lruCache *LRUCache[Key, Value]) Len() int {
	lruCache.mu.RLock()
	defer lruCache.mu.RUnlock()
	return lruCache.list.Len()
}
