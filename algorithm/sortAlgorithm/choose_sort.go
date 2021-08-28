package sortAlgorithm

func chooseSort(arr []int) []int {
	leng := len(arr)

	for i := 0; i < leng-1; i++ {
		min := i

		for j := i + 1; j < leng; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}

		swap(arr, i, min)
	}

	return arr
}
