package main

// 面试题3（一）：找出数组中重复的数字
// 题目：在一个长度为n的数组里的所有数字都在0到n-1的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，
// 也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。例如，如果输入长度为7的数组{2, 3, 1, 0, 2, 5, 3}，
// 那么对应的输出是重复的数字2或者3。

func duplicate(arr []int) (int, bool) {
	if len(arr) <= 1 {
		return -1, false
	}

	for _, v := range arr {
		if v < 0 || v > len(arr)-1 {
			// invalid arr
			return -1, false
		}
	}

	for i := 0; i < len(arr); i++ {
		for i != arr[i] {
			if arr[i] == arr[arr[i]] {
				return arr[i], true
			}

			arr[i], arr[arr[i]] = arr[arr[i]], arr[i]
		}
	}

	return -1, false
}
