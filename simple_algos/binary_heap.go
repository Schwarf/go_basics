package simple_algos

func swap[T any](a, b T) {
	a, b = b, a
}

type MinHeap struct {
	slice []int
}

func (heap *MinHeap) Insert(key int) {
	heap.slice = append(heap.slice, key)
	heap.heapifyUp(len(heap.slice) - 1)
}

func (heap *MinHeap) heapifyUp(index int) {
	for heap.slice[parent(index)] > heap.slice[index] {

		swap(parent(index), index)
		index = parent(index)
	}
}

func parent(index int) int {
	return (index - 1) / 2
}
