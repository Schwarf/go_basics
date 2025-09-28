package simple_algos

import "math/bits"

type Hasher[K comparable] func(K) uint64

type ShardedLRUCache[K comparable, V any] struct {
	shards []*LRUCache[K, V]
	mask   uint64
	hash   Hasher[K]
}

func NewShardedLRUCache[K comparable, V any](totalCapacity int, shardCount int, h Hasher[K]) *ShardedLRUCache[K, V] {
	if totalCapacity <= 0 {
		panic("ShardedLRU: totalCapacity must be > 0")
	}
	if shardCount <= 0 {
		panic("ShardedLRU: shardCount must be > 0")
	}
	if h == nil {
		panic("ShardedLRU: hash must not be nil")
	}
	perShard := totalCapacity / shardCount
	remain := totalCapacity % shardCount

	shards := make([]*LRUCache[K, V], shardCount)

	for i := 0; i < shardCount; i++ {
		capacity := perShard
		if i < remain {
			capacity++
		}
		shards[i] = NewLRUCache[K, V](capacity)
	}
	var mask uint64
	if 1<<bits.Len(uint(shardCount-1)) == shardCount {
		mask = uint64(shardCount - 1)
	}
	return &ShardedLRUCache[K, V]{shards: shards, hash: h, mask: mask}
}

func (s *ShardedLRUCache[K, V]) shardFor(key K) *LRUCache[K, V] {
	h := s.hash(key)
	if s.mask != 0 {
		return s.shards[int(h&s.mask)]
	}
	return s.shards[int(h%uint64(len(s.shards)))]
}

// Put inserts/updates a key in its shard.
func (s *ShardedLRUCache[K, V]) Put(key K, value V) {
	s.shardFor(key).Put(key, value)
}

// Get fetches and marks as most-recent within the shard.
func (s *ShardedLRUCache[K, V]) Get(key K) (V, bool) {
	return s.shardFor(key).Get(key)
}

// Peek fetches without affecting recency.
func (s *ShardedLRUCache[K, V]) Peek(key K) (V, bool) {
	return s.shardFor(key).Peek(key)
}

// Remove deletes key from its shard.
func (s *ShardedLRUCache[K, V]) Remove(key K) {
	s.shardFor(key).Remove(key)
}

// Len is the sum of shard lengths (point-in-time snapshot).
func (s *ShardedLRUCache[K, V]) Len() int {
	total := 0
	for _, sh := range s.shards {
		total += sh.Len()
	}
	return total
}

// Capacity returns the total capacity across all shards.
func (s *ShardedLRUCache[K, V]) Capacity() int {
	total := 0
	for _, sh := range s.shards {
		total += sh.Capacity()
	}
	return total
}
