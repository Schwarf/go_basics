package test_simple_algos

import (
	"testing"

	"github.com/Schwarf/go_basics/simple_algos"
)

func TestInsertAndRemoveBinaryHeap(t *testing.T) {
	heap := simple_algos.NewBinaryHeap[int](simple_algos.Less[int]{})
	valuesToInsert := []int{5, 3, 8, 1, 2}

	for _, val := range valuesToInsert {
		heap.Insert(val)
	}

	expectedOrder := []int{1, 2, 3, 5, 8}
	for _, expected := range expectedOrder {
		if value, error := heap.Pop(); error != nil || expected != value {
			t.Errorf("Heap did not remove elements in expected order.")
		}
	}
}

func TestInsertAndRemove(t *testing.T) {
	heap := simple_algos.MinHeap{}
	valuesToInsert := []int{5, 3, 8, 1, 2}

	for _, val := range valuesToInsert {
		heap.Insert(val)
	}

	expectedOrder := []int{1, 2, 3, 5, 8}
	for _, expected := range expectedOrder {
		if value, error := heap.Pop(); error != nil || expected != value {
			t.Errorf("Heap did not remove elements in expected order.")
		}
	}
}

func TestRemoveEmpty(t *testing.T) {
	heap := simple_algos.MinHeap{}
	if value, error := heap.Pop(); error == nil || value != 0 {
		t.Errorf("Empty heap does not return -1.")
	}
}
