package main

// 面试题53（二）：0到n-1中缺失的数字
// 题目：一个长度为n-1的递增排序数组中的所有数字都是唯一的，并且每个数字
// 都在范围0到n-1之内。在范围0到n-1的n个数字中有且只有一个数字不在该数组
// 中，请找出这个数字。

// 这个问题可以转换成在排序数组中找到第一个值和下标不相等的元素。

func getMissingNumber(arr []int) int {
	if len(arr) == 0 {
		return -1
	}

	start := 0
	end := len(arr) - 1

	for start <= end {
		mid := (start + end) / 2

		if arr[mid] != mid {
			if mid == 0 || arr[mid-1] == mid-1 {
				return mid
			} else {
				end = mid - 1
			}
		} else {
			start = mid + 1
		}
	}

	// 缺失最后一个数字
	if start == len(arr) {
		return start
	}

	return -1
}
