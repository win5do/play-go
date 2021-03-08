package sortAlgorithm

func heapSort(arr []int) []int {
	leng := len(arr)

	// 构建大顶堆
	buildMaxHeap(arr)

	// 将根节点(最大元素)与末尾元素互换，对剩余数组构建大顶堆
	for i := leng - 1; i > 0; i-- {
		swap(arr, 0, i)
		maxHeapify(arr, 0, i-1)
	}

	return arr
}

func buildMaxHeap(arr []int) {
	end := len(arr) - 1

	for i := (end - 1) / 2; i >= 0; i-- {
		maxHeapify(arr, i, end)
	}
}

func maxHeapify(arr []int, start, end int) {
	dad := start

	for 2*dad+1 <= end {
		l := 2*dad + 1
		r := 2*dad + 2
		large := dad

		if arr[l] > arr[dad] {
			large = l
		}

		if r <= end && arr[r] > arr[large] {
			large = r
		}

		if large != dad {
			swap(arr, dad, large)
			dad = large
		} else {
			return
		}
	}
}
