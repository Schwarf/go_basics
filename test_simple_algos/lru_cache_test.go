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
