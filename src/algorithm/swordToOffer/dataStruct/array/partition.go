package array

import (
	"math/rand"
)

func Partition(arr []int, start, end int) int {
	return partition_Hoare(arr, start, end)
}

func QuickSort(arr []int, start, end int) {
	if start >= end {
		return
	}

	index := Partition(arr, start, end)
	if index > start {
		QuickSort(arr, start, index)
	}
	if index < end {
		QuickSort(arr, index+1, end)
	}
}

// 两种partition算法对比，Hoare性能更好
// https://cs.stackexchange.com/questions/11458/quicksort-partitioning-hoare-vs-lomuto/11550

// Hoare
func partition_Hoare(arr []int, start, end int) int {
	if len(arr) == 0 || start < 0 || end >= len(arr) {
		panic("invalid input")
	}

	pivot := arr[end]

	left := start
	right := end

	for left < right {
		for left < right && arr[left] <= pivot {
			left++
		}

		for left < right && arr[right] >= pivot {
			right--
		}

		arr[left], arr[right] = arr[right], arr[left]
	}

	// 最后left == right，left多走一步到right的位置
	// 如果是取begin则应该right--放在前面
	arr[left], arr[end] = arr[end], arr[left]

	return right
}

// Lomuto
func partition_Lomuto(arr []int, start, end int) int {
	if len(arr) == 0 || start < 0 || end >= len(arr) {
		panic("invalid input")
	}

	// 随机选一个下标作为pivot，将其暂存到末尾
	random := rand.Intn(end-start) + start
	arr[end], arr[random] = arr[random], arr[end]
	// 可以跳过random直接取end
	pivot := arr[end]

	small := start
	for i := start; i < end; i++ {
		if arr[i] < pivot {
			if small != i {
				// 相等则不需要swap
				arr[i], arr[small] = arr[small], arr[i]
			}
			small++
		}
	}

	// 将mid放到分割线
	arr[small], arr[end] = arr[end], arr[small]
	return small
}
