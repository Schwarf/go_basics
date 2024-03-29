package test_simple_algos

import (
	"testing"

	"github.com/Schwarf/go_basics/simple_algos"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		slice    []int
		target   int
		expected int
	}{
		{[]int{1, 2, 3, 4, 5}, 3, 2},
		{[]int{1, 2, 3, 4, 5}, 1, 0},
		{[]int{1, 2, 3, 4, 5}, 5, 4},
		{[]int{1, 2, 3, 4, 5}, 6, -1},
		{[]int{10, 20, 30, 40, 50}, 40, 3},
	}

	for _, tt := range tests {
		result := simple_algos.BinarySearch(tt.slice, tt.target)
		if result != tt.expected {
			t.Errorf("BinarySearch(%v, %d) = %d; want %d", tt.slice, tt.target, result, tt.expected)
		}
	}
}
