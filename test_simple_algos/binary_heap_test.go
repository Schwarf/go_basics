package test_simple_algos

import (
	"testing"

	"github.com/Schwarf/go_basics/simple_algos"
)

func TestInsertAndRemove(t *testing.T) {
	heap := simple_algos.MinHeap{}
	valuesToInsert := []int{5, 3, 8, 1, 2}

	for _, val := range valuesToInsert {
		heap.Insert(val)
	}

	expectedOrder := []int{1, 2, 3, 5, 8}
	for _, expected := range expectedOrder {
		if heap.Pop() != expected {
			t.Errorf("Heap did not remove elements in expected order.")
		}
	}
}
