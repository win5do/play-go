package main

// 面试题53（三）：数组中数值和下标相等的元素
// 题目：假设一个单调递增的数组里的每个元素都是整数并且是唯一的。请编程实
// 现一个函数找出数组中任意一个数值等于其下标的元素。例如，在数组{-3, -1,
// 1, 3, 5}中，数字3和它的下标相等。

func getNumberEqualIndex(arr []int) int {
	if len(arr) == 0 {
		return -1
	}

	start := 0
	end := len(arr) - 1
	mid := (start + end) >> 1
	for start <= end {
		if arr[mid] == mid {
			return mid
		} else if arr[mid] > mid {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return -1
}
