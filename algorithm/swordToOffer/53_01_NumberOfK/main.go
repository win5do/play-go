package main

// 面试题53（一）：数字在排序数组中出现的次数
// 题目：统计一个数字在排序数组中出现的次数。例如输入排序数组{1, 2, 3, 3,
// 3, 3, 4, 5}和数字3，由于3在这个数组中出现了4次，因此输出4。

func getFirstK(arr []int, k, start, end int) int {
	if len(arr) == 0 || start < 0 || end > len(arr)-1 || start > end {
		return -1
	}

	mid := (start + end) / 2
	if arr[mid] == k {
		if mid == 0 || mid > 0 && arr[mid-1] != k {
			return mid
		} else {
			end = mid - 1
		}
	} else if arr[mid] > k {
		end = mid - 1
	} else {
		start = mid + 1
	}

	return getFirstK(arr, k, start, end)
}

func getLastK(arr []int, k, start, end int) int {
	length := len(arr)

	if length == 0 || start < 0 || end > len(arr)-1 || start > end {
		return -1
	}

	mid := (start + end) / 2
	if arr[mid] == k {
		if mid == length-1 || mid < length-1 && arr[mid+1] != k {
			return mid
		} else {
			start = mid + 1
		}
	} else if arr[mid] > k {
		end = mid - 1
	} else {
		start = mid + 1
	}

	return getLastK(arr, k, start, end)
}

func getNumberOfK(arr []int, k int) int {
	if len(arr) == 0 {
		return 0
	}

	first := getFirstK(arr, k, 0, len(arr)-1)
	if first == -1 {
		return 0
	}

	last := getLastK(arr, k, first, len(arr)-1)
	if last == -1 {
		return 0
	}

	return last - first + 1
}
