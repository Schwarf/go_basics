package test_simple_algos

import (
	"testing"

	"github.com/Schwarf/go_basics/simple_algos"
)

func TestMergeSearch(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{2, 1, 5, 4, 3}, []int{1, 2, 3, 4, 5}},
		{[]int{2, -1, -5, 4, 3}, []int{-5, -1, 2, 3, 4}},
	}

	for _, tt := range tests {
		result := simple_algos.MergeSort(tt.input)
		if result != tt.expected {
			t.Errorf("MergeSort(%v) != %v;", tt.input, tt.expected)
		}

	}

}
