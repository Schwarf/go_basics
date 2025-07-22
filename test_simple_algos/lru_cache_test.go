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

func TestLRUCacheMostRecent(t *testing.T) {
	length := 3
	cache := simple_algos.NewLRUCache[int, int](length)
	cache.Put(1, 1)

}
