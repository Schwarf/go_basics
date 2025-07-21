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
