package simple_algos

import "fmt"

type MinHeap struct {
	slice []int
}

func (heap *MinHeap) swap(i, j int) {
	heap.slice[i], heap.slice[j] = heap.slice[j], heap.slice[i]
}

func (heap *MinHeap) Insert(key int) {
	heap.slice = append(heap.slice, key)
	heap.heapifyUp(len(heap.slice) - 1)
}

func (heap *MinHeap) heapifyUp(index int) {
	for heap.slice[parent(index)] > heap.slice[index] {
		heap.swap(parent(index), index)
		index = parent(index)
	}
}

func parent(index int) int {
	return (index - 1) / 2
}

func (heap *MinHeap) Pop() int {
	if len(heap.slice) == 0 {
		fmt.Println("Heap is empty")
		return -1
	}

	minimum := heap.slice[0]
	heap.slice[0] = heap.slice[len(heap.slice)-1]
	heap.slice = heap.slice[:len(heap.slice)-1]
	heap.heapifyDown(0)
	return minimum
}

func (heap *MinHeap) heapifyDown(index int) {
	lastIndex := len(heap.slice) - 1
	leftIndex, rightIndex := leftChild(index), rightChild(index)
	indexToCompare := 0

	for leftIndex <= lastIndex {
		if leftIndex == lastIndex { // When left child is the only child
			indexToCompare = leftIndex
		} else if heap.slice[leftIndex] < heap.slice[rightIndex] {
			indexToCompare = leftIndex
		} else {
			indexToCompare = rightIndex
		}

		if heap.slice[index] > heap.slice[indexToCompare] {
			heap.swap(index, indexToCompare)
			index = indexToCompare
			leftIndex, rightIndex = leftChild(index), rightChild(index)
		} else {
			return
		}
	}
}

func leftChild(index int) int {
	return 2*index + 1
}

func rightChild(index int) int {
	return 2*index + 2
}
