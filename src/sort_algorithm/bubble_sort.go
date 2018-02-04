package sort_algorithm

func bubbleSort(arr []int) []int {
	l := len(arr)
	for i := 0; i < l-1; i++ {
		for j := 0; j < l-1-i; j++ {
			if arr[j] > arr[j+1] {
				t := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = t
			}
		}
	}
	return arr
}
