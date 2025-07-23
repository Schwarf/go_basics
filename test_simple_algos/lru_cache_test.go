package test_simple_algos

import (
	"github.com/Schwarf/go_basics/simple_algos"
	"testing"
)

func TestLRUCacheLen(t *testing.T) {
	length := 3
	cache := simple_algos.NewLRUCache[int, int](length)
	valuesToInsert := []int{5, 3, 8, 1, 2}

	for index, val := range valuesToInsert {
		cache.Put(index, val)
	}

	if actualLen := cache.Len(); actualLen != length {
		t.Errorf("Expected cache length %d, got %d", length, actualLen)
	}
}

func TestLRUCachePeek(t *testing.T) {
	length := 3
	cache := simple_algos.NewLRUCache[int, int](length)
	cache.Put(1, 1)

	_, notOk := cache.Peek(2)
	if notOk {
		t.Errorf("Found an expected element")
	}

	_, ok := cache.Peek(1)
	if !ok {
		t.Errorf("Did not find an expected element")
	}
}

func TestLRUCacheEvictPolicy(t *testing.T) {
	length := 3
	cache := simple_algos.NewLRUCache[int, int](length)
	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Put(3, 3)
	_, ok := cache.Peek(1)
	if !ok {
		t.Errorf("Did not find an expected element")
	}

	cache.Put(4, 4)

	_, notOk := cache.Peek(1)
	if notOk {
		t.Errorf("Found an expected element")
	}
}

func TestLRUCacheGet(t *testing.T) {
	// capacity = 2
	cache := simple_algos.NewLRUCache[int, int](2)

	// Getting a non‑existent key should return zero value + false
	if v, ok := cache.Get(42); ok {
		t.Errorf("Get on empty cache: expected ok=false, got ok=true with value %v", v)
	}

	// Insert two items
	cache.Put(1, 100)
	cache.Put(2, 200)

	// Get(1) should return 100, true
	if v, ok := cache.Get(1); !ok || v != 100 {
		t.Errorf("Get(1): expected (100, true), got (%v, %v)", v, ok)
	}

	// Now that 1 was accessed, it is the most recent.
	// Insert a third item, evicting the least recent (which is key=2)
	cache.Put(3, 300)

	// Key 2 must have been evicted
	if _, ok := cache.Get(2); ok {
		t.Errorf("After Put(3): expected key 2 to be evicted, but Get(2) returned ok=true")
	}

	// And key 1 should still be present
	if v, ok := cache.Get(1); !ok || v != 100 {
		t.Errorf("After eviction: expected Get(1) to still return (100, true), got (%v, %v)", v, ok)
	}

	// Finally, Get(3) should work
	if v, ok := cache.Get(3); !ok || v != 300 {
		t.Errorf("Get(3): expected (300, true), got (%v, %v)", v, ok)
	}
}
