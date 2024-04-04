package test_simple_algos

import (
	"testing"

	"github.com/Schwarf/go_basics/simple_algos"
)

func TestInsertAndRemoveMinHeap(t *testing.T) {
	minHeap := simple_algos.NewBinaryHeap[int](simple_algos.Less[int]{})
	valuesToInsert := []int{5, 3, 8, 1, 2}

	for _, val := range valuesToInsert {
		minHeap.Insert(val)
	}

	expectedOrder := []int{1, 2, 3, 5, 8}
	for _, expected := range expectedOrder {
		if value, error := minHeap.Pop(); error != nil || expected != value {
			t.Errorf("Heap did not remove elements in expected order.")
		}
	}
}

func TestInsertAndRemoveMaxHeap(t *testing.T) {
	maxHeap := simple_algos.NewBinaryHeap[float64](simple_algos.Greater[float64]{})
	valuesToInsert := []float64{5, 3, 8, 1, 2}

	for _, val := range valuesToInsert {
		maxHeap.Insert(val)
	}

	expectedOrder := []float64{8, 5, 3, 2, 1}
	for _, expected := range expectedOrder {
		if value, error := maxHeap.Pop(); error != nil || expected != value {
			t.Errorf("Heap did not remove elements in expected order. Expected: %.2f, Is: %.2f", expected, value)
		}
	}
}

func TestRemoveEmptyMinHeap(t *testing.T) {
	heap := simple_algos.NewBinaryHeap[int](simple_algos.Less[int]{})
	if value, error := heap.Pop(); error == nil || value != 0 {
		t.Errorf("Empty heap does not return -1.")
	}
}

func TestPopMinHeap(t *testing.T) {
	minHeap := simple_algos.NewBinaryHeap[int8](simple_algos.Less[int8]{})
	valuesToInsert := []int8{5, 3, 8, 1, 2}
	for _, val := range valuesToInsert {
		minHeap.Insert(val)
	}

	expectedOrder := []int8{1, 2, 3, 5, 8}
	for _, expected := range expectedOrder {
		if value, error := minHeap.Top(); error != nil || expected != value {
			t.Errorf("Heap did not show elements in expected order. Expected: %.2d, Is: %.2d", expected, value)
		}
		if value, error := minHeap.Pop(); error != nil || expected != value {
			t.Errorf("Heap did not remove elements in expected order. Expected: %.2d, Is: %.2d", expected, value)
		}

	}

}
