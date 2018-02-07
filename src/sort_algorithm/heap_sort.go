package sort_algorithm

import "fmt"

func heapSort(arr []int) []int {
	buildMaxHeap(arr)

	fmt.Println(arr)

	for i := len(arr) - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		buildMaxHeap(arr[:i])
	}

	return arr
}

func buildMaxHeap(arr []int) {
	leng := len(arr)
	if leng <= 1 {
		return
	}

	for i := (leng - 1) / 2; i >= 0; i-- {
		max := i
		left := 2*i + 1
		right := 2*i + 2

		if left < leng && arr[max] < arr[left] {
			max = left
		}

		if right < leng && arr[max] < arr[right] {
			max = right
		}

		if i != max {
			arr[i], arr[max] = arr[max], arr[i]
		}
	}
}
