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

	var extremum T
	if len(heap.elements) == 0 {
		return extremum, errors.New("BINARY HEAP IS EMPTY")
	}

	extremum = heap.elements[0]
	heap.elements[0] = heap.elements[len(heap.elements)-1]
	heap.elements = heap.elements[:len(heap.elements)-1]
	heap.heapifyDown()
	return extremum, nil

}

func (heap *BinaryHeap[T]) Top() (T, error) {

	var extremum T
	if len(heap.elements) == 0 {
		return extremum, errors.New("BINARY HEAP IS EMPTY")
	}
	extremum = heap.elements[0]
	return extremum, nil

}

func (heap *BinaryHeap[T]) heapifyDown() {
	// we start from the top
	lastIndex := len(heap.elements) - 1
	index := 0
	leftIndex, rightIndex := 1, 2
	childIndexToCompare := 0

	for leftIndex <= lastIndex {

		if leftIndex == lastIndex { // When left child is the only child
			childIndexToCompare = leftIndex
		} else if heap.comparator.Compare(heap.elements[leftIndex], heap.elements[rightIndex]) { // When left child is less/greater (MIN/MAX CASE)than right child
			childIndexToCompare = leftIndex
		} else { // When right child is less/greater than left child
			childIndexToCompare = rightIndex
		}

		if heap.comparator.Compare(heap.elements[childIndexToCompare], heap.elements[index]) { // When child is less/greater (min/max case) than element, swap ...
			heap.swap(index, childIndexToCompare)
			index = childIndexToCompare
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
	heap.heapifyUp()
}

func (heap *BinaryHeap[T]) heapifyUp() {
	// When element is less/greater than the parent move it up
	index := len(heap.elements) - 1
	for heap.comparator.Compare(heap.elements[index], heap.elements[heap.parent(index)]) {
		heap.swap(parent(index), index)
		index = parent(index)
	}
}

func parent(index int) int {
	return (index - 1) / 2
}

func leftChild(index int) int {
	return 2*index + 1
}

func rightChild(index int) int {
	return 2*index + 2
}
