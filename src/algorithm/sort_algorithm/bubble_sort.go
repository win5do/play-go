package sort_algorithm

func bubbleSort(arr []int) []int {
	leng := len(arr)
	for i := leng - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				swap(arr, j, j+1)
			}
		}
	}
	return arr
}
