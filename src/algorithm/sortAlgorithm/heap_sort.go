package sortAlgorithm

func heapSort(arr []int) []int {
	l := len(arr)

	// 构建大顶堆
	for i := l/2 - 1; i >= 0; i-- {
		adjustHeap(arr, i)
	}

	// 将根节点(最大元素)与末尾元素互换，对剩余数组构建大顶堆
	for i := l - 1; i > 0; i-- {
		swap(arr, 0, i)
		adjustHeap(arr[:i], 0)
	}

	return arr
}

// 如果arr[k]比左右子节点小则下沉，直至叶子节点
func adjustHeap(arr []int, k int) {
	l := len(arr)

	temp := arr[k]

	for i := 2*k + 1; i < l; i = 2*i + 1 {
		if i+1 < l && arr[i+1] > arr[i] {
			i++
		}

		if (arr[i]) > temp {
			arr[k] = arr[i]
			k = i
		} else {
			break
		}
	}

	arr[k] = temp
}
