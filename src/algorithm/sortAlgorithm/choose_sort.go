package sortAlgorithm

func chooseSort(arr []int) []int {
	leng := len(arr)

	for i := 0; i < leng-1; i++ {
		min := arr[i]

		for j := i + 1; j < leng; j++ {
			if arr[j] < min {
				t := min
				min = arr[j]
				arr[j] = t
			}
		}

		arr[i] = min
	}

	return arr
}
