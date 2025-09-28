package simple_algos

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

}
