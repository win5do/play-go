package sortAlgorithm

func insertSort(arr []int) []int {
	leng := len(arr)

	for i := 1; i < leng; i++ {
		temp := arr[i]
		j := i - 1

		for j >= 0 && temp < arr[j] {
			arr[j+1] = arr[j]
			j--
		}

		arr[j+1] = temp
	}

	return arr
}
