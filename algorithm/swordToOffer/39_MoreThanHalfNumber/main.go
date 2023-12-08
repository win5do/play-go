package main

import (
	"play-go/algorithm/swordToOffer/dataStruct/array"
)

// 面试题39：数组中出现次数超过一半的数字
// 题目：数组中有一个数字出现的次数超过数组长度的一半，请找出这个数字。例
// 如输入一个长度为9的数组{1, 2, 3, 2, 2, 2, 5, 4, 2}。由于数字2在数组中
// 出现了5次，超过数组长度的一半，因此输出2。

func checkMoreThanHalf(arr []int, n int) bool {
	count := 0
	for _, v := range arr {
		if v == n {
			count++
		}
	}

	// 如果用arr/2还有判断奇偶，所以使用count*2
	return count<<1 >= len(arr)
}

// --- partition solution ---
func moreThanHalfNum_partition(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	mid := len(arr) >> 1
	start := 0
	end := len(arr) - 1
	index := array.Partition(arr, start, end)

	// 如果选中的数字的下标刚好是n/2，那么这个数字就是中位数
	// 如果下标大于n/2，则中位数应该在塔左边；反之则在右边
	for index != mid {
		if index > mid {
			end = index - 1
			index = array.Partition(arr, start, end)
		} else {
			start = index + 1
			index = array.Partition(arr, start, end)
		}
	}

	result := arr[index]
	if checkMoreThanHalf(arr, result) {
		return result
	}

	return 0
}

// --- count solution ---
func moreThanHalfNum_count(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	result := arr[0]
	count := 1
	for i := 1; i < len(arr); i++ {
		if count == 0 {
			result = arr[i]
			count = 1
		} else if arr[i] == result {
			count++
		} else {
			count--
		}
	}

	if checkMoreThanHalf(arr, result) {
		return result
	}

	return 0
}
