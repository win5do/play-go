package sortAlgorithm

func bubbleSort(arr []int) []int {
	leng := len(arr)

	for i := leng - 1; i > 0; i-- {
		unchange := true
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				swap(arr, j, j+1)
				unchange = false
			}
		}

		if unchange {
			// 如果有一趟没有发生交换说明排序已经完成
			break
		}
	}

	return arr
}
