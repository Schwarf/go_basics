package test_simple_algos

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"

	"github.com/Schwarf/go_basics/simple_algos"
)

func TestNewLRUCacheNegativeCapacity(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected panic for negative capacity, got none")
		}
	}()

	_ = simple_algos.NewLRUCache[float64, int](-1)
}

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

func TestLRUCachePutUpdatesValue(t *testing.T) {
	cache := simple_algos.NewLRUCache[int, string](2)
	cache.Put(1, "one")
	cache.Put(1, "uno") // update existing key

	if got, ok := cache.Get(1); !ok || got != "uno" {
		t.Errorf("After updating key=1, expected (\"uno\", true), got (%q, %v)", got, ok)
	}
	if length := cache.Len(); length != 1 {
		t.Errorf("Expected length=1 after updating existing key, got %d", length)
	}
}

func TestLRUCacheEmpty(t *testing.T) {
	cache := simple_algos.NewLRUCache[int, int](1)
	if got := cache.Len(); got != 0 {
		t.Errorf("New cache.Len() = %d, want 0", got)
	}
	if _, ok := cache.Get(42); ok {
		t.Errorf("Get on empty cache returned ok=true; want ok=false")
	}
	if _, ok := cache.Peek(42); ok {
		t.Errorf("Peek on empty cache returned ok=true; want ok=false")
	}
}

func TestLRUCacheGetMovesToFront(t *testing.T) {
	cache := simple_algos.NewLRUCache[int, int](2)
	cache.Put(1, 10)
	cache.Put(2, 20)

	// Access key=1, making it most‑recent
	if v, ok := cache.Get(1); !ok || v != 10 {
		t.Fatalf("Get(1) = (%v, %v), want (10, true)", v, ok)
	}

	// Insert 3 → should evict the LRU (which is now key=2)
	cache.Put(3, 30)

	if _, ok := cache.Get(2); ok {
		t.Errorf("Expected key=2 to be evicted after Put(3), but Get(2) returned ok=true")
	}
	// Keys 1 and 3 should still be present
	for _, tc := range []struct{ key, want int }{{1, 10}, {3, 30}} {
		if v, ok := cache.Get(tc.key); !ok || v != tc.want {
			t.Errorf("Get(%d) = (%v, %v), want (%d, true)", tc.key, v, ok, tc.want)
		}
	}
}

func TestLRUCacheConcurrentAccess(t *testing.T) {
	cache := simple_algos.NewLRUCache[int, int](50) // capacity 50
	var wg sync.WaitGroup

	nWorkers := 8
	nOps := 10000000

	for w := 0; w < nWorkers; w++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for i := 0; i < nOps; i++ {
				key := rand.Intn(1000)
				val := id*1000 + i
				switch rand.Intn(3) {
				case 0:
					cache.Put(key, val)
				case 1:
					cache.Get(key)
				case 2:
					cache.Remove(key)
				}
			}
		}(w)
	}
	wg.Wait()

	if got := cache.Len(); got > 50 {
		t.Errorf("Len() = %d, exceeds capacity 50", got)
	}
}

func TestLRUCacheCapacity(t *testing.T) {
	cap := 5
	cache := simple_algos.NewLRUCache[int, string](cap)

	if got := cache.Capacity(); got != cap {
		t.Errorf("Capacity() = %d, want %d", got, cap)
	}

	// Fill it and ensure Len ≤ Capacity
	for i := 0; i < 10; i++ {
		cache.Put(i, fmt.Sprintf("val-%d", i))
	}
	if cache.Len() > cache.Capacity() {
		t.Errorf("Len() = %d exceeds Capacity() = %d", cache.Len(), cache.Capacity())
	}
}
