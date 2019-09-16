package main

// 面试题3（二）：不修改数组找出重复的数字
// 题目：在一个长度为n+1的数组里的所有数字都在1到n的范围内，所以数组中至
// 少有一个数字是重复的。请找出数组中任意一个重复的数字，但不能修改输入的
// 数组。例如，如果输入长度为8的数组{2, 3, 5, 4, 3, 2, 6, 7}，那么对应的
// 输出是重复的数字2或者3。

func duplicate(arr []int) (int, bool) {
	if len(arr) <= 1 {
		return -1, false
	}

	start := 1
	end := len(arr) - 1

	for end >= start {
		// equal (start + end) / 2
		mid := (start + end) >> 1
		count := countRange(arr, start, mid)
		if start == end {
			if count > 1 {
				return start, true
			} else {
				break
			}
		}

		if count > mid-start+1 {
			end = mid
		} else {
			start = mid + 1
		}
	}

	return -1, false
}

func countRange(arr []int, start, end int) int {
	count := 0
	for _, v := range arr {
		if v >= start && v <= end {
			count++
		}
	}

	return count
}
