package searchAlgorithm

func BinarySearch(arr []int, key int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		m := (left + right) / 2
		v := arr[m]
		if v == key {
			return m
		} else if v > key {
			right = m - 1
		} else {
			left = m + 1
		}
	}
	return -1
}
