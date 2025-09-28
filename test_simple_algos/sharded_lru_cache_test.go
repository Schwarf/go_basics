package test_simple_algos

import (
	"math/rand"
	"sync"
	"testing"

	"github.com/Schwarf/go_basics/simple_algos"
)

// Simple int hasher for tests.
func intHasher(x int) uint64 { return uint64(x) }

func TestShardedLRUBasicPutGet(t *testing.T) {
	slru := simple_algos.NewShardedLRUCache[int, string](10, 2, intHasher)

	slru.Put(1, "one")
	slru.Put(2, "two")

	if v, ok := slru.Get(1); !ok || v != "one" {
		t.Errorf("Get(1) = (%q, %v), want (\"one\", true)", v, ok)
	}
	if v, ok := slru.Peek(2); !ok || v != "two" {
		t.Errorf("Peek(2) = (%q, %v), want (\"two\", true)", v, ok)
	}

	slru.Remove(1)
	if _, ok := slru.Get(1); ok {
		t.Errorf("Expected key=1 to be removed")
	}
}

func TestShardedLRUCapacity(t *testing.T) {
	totalCap := 8
	slru := simple_algos.NewShardedLRUCache[int, int](totalCap, 4, intHasher)

	if got := slru.Capacity(); got != totalCap {
		t.Errorf("Capacity() = %d, want %d", got, totalCap)
	}

	// Insert more than capacity and check Len() ≤ Capacity
	for i := 0; i < 20; i++ {
		slru.Put(i, i)
	}
	if slru.Len() > slru.Capacity() {
		t.Errorf("Len() = %d exceeds Capacity() = %d", slru.Len(), slru.Capacity())
	}
}

func TestShardedLRUShardDistribution(t *testing.T) {
	// With 4 shards and identity hasher, keys 0,1,2,3 go to different shards
	slru := simple_algos.NewShardedLRUCache[int, int](8, 4, intHasher)
	for i := 0; i < 4; i++ {
		slru.Put(i, i)
	}

	// If shardFor were broken (all into shard 0), Len would equal capacity of shard 0
	if slru.Len() != 4 {
		t.Errorf("Expected 4 items across shards, got Len=%d", slru.Len())
	}
}

func TestShardedLRUConcurrent(t *testing.T) {
	slru := simple_algos.NewShardedLRUCache[int, int](100, 8, intHasher)

	var wg sync.WaitGroup
	nOps := 10000000
	for w := 0; w < 16; w++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for i := 0; i < nOps; i++ {
				key := rand.Intn(1000)
				slru.Put(key, i)
				slru.Get(key)
				slru.Remove(key)
			}
		}(w)
	}
	wg.Wait()

	if slru.Len() > slru.Capacity() {
		t.Errorf("Concurrent Len()=%d exceeds Capacity=%d", slru.Len(), slru.Capacity())
	}
}

func TestShardedLRUNotPowerOfTwoShards(t *testing.T) {
	slru := simple_algos.NewShardedLRUCache[int, int](10, 3, intHasher)

	for i := 0; i < 9; i++ {
		slru.Put(i, i)
	}
	if slru.Len() > slru.Capacity() {
		t.Errorf("Len=%d exceeds Capacity=%d", slru.Len(), slru.Capacity())
	}
}
