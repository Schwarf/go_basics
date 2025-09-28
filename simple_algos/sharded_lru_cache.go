package simple_algos

import "math/bits"

type Hasher[K comparable] func(K) uint64

type ShardedLRUCache[K comparable, V any] struct {
	shards []*LRUCache[K, V]
	mask   uint64
	hasher Hasher[K]
}

func newShardedLRUCache[K comparable, V any](totalCapacity int, shardCount int, h Hasher[K]) *ShardedLRUCache[K, V] {
	if totalCapacity <= 0 {
		panic("ShardedLRU: totalCapacity must be > 0")
	}
	if shardCount <= 0 {
		panic("ShardedLRU: shardCount must be > 0")
	}
	if h == nil {
		panic("ShardedLRU: hasher must not be nil")
	}
	perShard := totalCapacity / shardCount
	remain := totalCapacity % shardCount

	shards := make([]*LRUCache[K, V], shardCount)

	for i := 0; i < shardCount; i++ {
		capacity := perShard
		if i == shardCount-1 {
			capacity = remain
		}
		if capacity <= 0 {
			capacity = 1
		}
		shards[i] = NewLRUCache[K, V](capacity)
	}
	var mask uint64
	if 1<<bits.Len(uint64(shardCount-1)) == shardCount {
		mask = uint64(shardCount - 1)
	}
	return &ShardedLRUCache[K, V]{shards: shards, hasher: h, mask: mask}
}

func (s *ShardedLRUCache[K, V]) shardFor(key K) *LRUCache[K, V] {
	h := s.hasher(key)
	if s.mask != 0 {
		return s.shards[int(h&s.mask)]
	}
	return s.shards[int(h%uint64(len(s.shards)))]
}
