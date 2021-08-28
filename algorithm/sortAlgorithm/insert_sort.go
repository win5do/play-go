package sortAlgorithm

func insertSort(arr []int) []int {
	leng := len(arr)

	for i := 1; i < leng; i++ {
		tmp := arr[i]
		j := i - 1

		for ; j >= 0 && tmp < arr[j]; j-- {
			arr[j+1] = arr[j]
		}

		arr[j+1] = tmp
	}

	return arr
}
