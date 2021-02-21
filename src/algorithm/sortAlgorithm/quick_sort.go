package sortAlgorithm

import (
	"math/rand"
)

func QuickSort_Lomuto(arr []int, start, end int) {
	if start >= end {
		return
	}

	index := Partition_Lomuto(arr, start, end)
	QuickSort_Lomuto(arr, start, index-1)
	QuickSort_Lomuto(arr, index+1, end)
}

func QuickSort_Hoare(arr []int, start, end int) {
	if start >= end {
		return
	}

	index := Partition_Lomuto(arr, start, end)
	QuickSort_Hoare(arr, start, index)
	QuickSort_Hoare(arr, index+1, end)
}

// 两种 partition 算法对比，Hoare 性能更好，Lomuto 实现简单
// https://cs.stackexchange.com/questions/11458/quicksort-partitioning-hoare-vs-lomuto/11550

// https://en.wikipedia.org/wiki/Quicksort#Lomuto_partition_scheme
func Partition_Lomuto(arr []int, start, end int) int {
	if start >= end {
		return start
	}

	// 随机选一个下标作为pivot，将其暂存到末尾
	random := rand.Intn(end-start) + start
	arr[end], arr[random] = arr[random], arr[end]
	// 可以跳过random直接取end
	pivot := arr[end]

	l := start
	for i := start; i < end; i++ {
		if arr[i] < pivot {
			swap(arr, i, l)
			l++
		}
	}

	// 将mid放到分割线
	swap(arr, end, l)
	return l
}

// https://en.wikipedia.org/wiki/Quicksort#Hoare_partition_scheme
func Partition_Hoare(arr []int, start, end int) int {
	if start >= end {
		return start
	}

	pivot := arr[(start+end)/2]

	left := start - 1
	right := end + 1

	for {
		left++
		for arr[left] < pivot {
			left++
		}

		right--
		for arr[right] > pivot {
			right--
		}

		if left >= right {
			return right
		}

		swap(arr, left, right)
	}
}
