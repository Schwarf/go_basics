package test_simple_algos

import (
	"fmt"
	"testing"

	"math/rand"
	"sort"

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

func TestMaxHeapWithRandomNumbers(t *testing.T) {
	randomNumbers := make([]int, 1000)
	maxHeap := simple_algos.NewBinaryHeap[int](simple_algos.Greater[int]{})
	for i := range randomNumbers {
		randomNumbers[i] = rand.Intn(10000)
		maxHeap.Insert(int(randomNumbers[i]))
	}
	sort.Sort(sort.Reverse(sort.IntSlice(randomNumbers)))
	for _, expected := range randomNumbers {
		extracted, err := maxHeap.Pop() // Implement Pop according to your heap
		if err != nil {
			fmt.Println("Error extracting from heap:", err)
			break
		}
		if extracted != expected {
			fmt.Println("Mismatch! Expected:", expected, "Got:", extracted)
			break
		}
	}

}
