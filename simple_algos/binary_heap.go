package simple_algos

import (
	"errors"

	"golang.org/x/exp/constraints"
)

type Comparator[T constraints.Ordered] interface {
	Compare(a, b T) bool
}

type Less[T constraints.Ordered] struct{}

func (Less[T]) Compare(a, b T) bool {
	return a < b
}

type Greater[T constraints.Ordered] struct{}

func (Greater[T]) Compare(a, b T) bool {
	return a > b
}

type BinaryHeap[T constraints.Ordered] struct {
	elements   []T
	comparator Comparator[T]
}

func NewBinaryHeap[T constraints.Ordered](compare Comparator[T]) *BinaryHeap[T] {
	return &BinaryHeap[T]{
		elements:   make([]T, 0),
		comparator: compare,
	}
}

func (heap *BinaryHeap[T]) parent(index int) int {
	return (index - 1) / 2
}

func (heap *BinaryHeap[T]) Pop() (T, error) {

	var minimum T
	if len(heap.elements) == 0 {
		return minimum, errors.New("Binary heap is empty")
	}

	minimum = heap.elements[0]
	heap.elements[0] = heap.elements[len(heap.elements)-1]
	heap.elements = heap.elements[:len(heap.elements)-1]
	heap.heapifyDown(0)
	return minimum, nil

}

func (heap *BinaryHeap[T]) heapifyDown(index int) {
	lastIndex := len(heap.elements) - 1
	leftIndex, rightIndex := leftChild(index), rightChild(index)
	indexToCompare := 0

	for leftIndex <= lastIndex {
		if leftIndex == lastIndex { // When left child is the only child
			indexToCompare = leftIndex
		} else if heap.elements[leftIndex] < heap.elements[rightIndex] {
			indexToCompare = leftIndex
		} else {
			indexToCompare = rightIndex
		}

		if heap.comparator.Compare(heap.elements[indexToCompare], heap.elements[index]) {
			heap.swap(index, indexToCompare)
			index = indexToCompare
			leftIndex, rightIndex = leftChild(index), rightChild(index)
		} else {
			return
		}
	}

}

func (heap *BinaryHeap[T]) swap(i, j int) {
	heap.elements[i], heap.elements[j] = heap.elements[j], heap.elements[i]
}

func (heap *BinaryHeap[T]) Insert(key T) {
	heap.elements = append(heap.elements, key)
	heap.heapifyUp(len(heap.elements) - 1)
}

func (heap *BinaryHeap[T]) heapifyUp(index int) {
	for heap.comparator.Compare(heap.elements[index], heap.elements[heap.parent(index)]) {
		heap.swap(parent(index), index)
		index = parent(index)
	}
}

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

func (heap *MinHeap) Pop() (int, error) {
	if len(heap.slice) == 0 {
		return 0, errors.New("MinHeap is empty")
	}

	minimum := heap.slice[0]
	heap.slice[0] = heap.slice[len(heap.slice)-1]
	heap.slice = heap.slice[:len(heap.slice)-1]
	heap.heapifyDown(0)
	return minimum, nil
}

func (heap *MinHeap) Top() (int, error) {
	if len(heap.slice) == 0 {

		return 0, errors.New("MinHeap is empty")
	}
	return heap.slice[0], nil
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
