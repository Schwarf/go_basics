package simple_algos

func BinarySearch(slice []int, target int) int {
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
