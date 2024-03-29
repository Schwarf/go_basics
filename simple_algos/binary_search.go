package simple_algos

import "golang.org/x/exp/constraints"

func BinarySearch[T constraints.Ordered](slice []T, target T) int {
	left := 0
	right := len(slice) - 1
	for left <= right {
		mid := (left + right) / 2

		if slice[mid] == target {
			return mid
		} else if slice[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
