package simple_algos

import "container/list"

type LRUCache[Key comparable, Value any] struct {
	capacity int
	cache    map[Key]*list.Element
	list     *list.List
}

type lruCacheEntry[Key comparable, Value any] struct {
	key   Key
	value Value
}

func New[Key comparable, Value any](capacity int) *LRUCache[Key, Value] {
	return &LRUCache[Key, Value]{
		capacity: capacity,
		cache:    make(map[Key]*list.Element),
		list:     list.New()}
}

func (lruCache *LRUCache[Key, Value]) Get(key Key) (Value, bool) {
	if element, ok := lruCache.cache[key]; ok {
		lruCache.list.MoveToFront(element)
		return element.Value.(lruCacheEntry[Key, Value]).value, true
	}
	var zero Value // is nil (where nil is allowed) and 0 (false) for others like int(bool)
	return zero, false
}
