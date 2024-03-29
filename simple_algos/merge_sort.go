package simple_algos

import "golang.org/x/exp/constraints"

func merge[T constraints.Ordered](leftArray, rightArray []T) []T {

	var result []T
	left, right := 0, 0
	for left < len(leftArray) && right < len(rightArray) {
		if leftArray[left] < rightArray[right] {
			result = append(result, leftArray[left])
			left++
		} else {
			result = append(result, rightArray[right])
			right++
		}
	}
	result = append(result, leftArray[left:]...)
	result = append(result, rightArray[right:]...)
	return result
}

func MergeSort[T constraints.Ordered](input []T) []T {
	if len(input) < 2 {
		return input
	}

	mid := len(input) / 2
	leftArray := MergeSort(input[:mid])
	rightArray := MergeSort(input[mid:])
	return merge(leftArray, rightArray)
}
