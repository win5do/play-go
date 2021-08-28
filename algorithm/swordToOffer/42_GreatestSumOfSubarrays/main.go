package main

// 面试题42：连续子数组的最大和
// 题目：输入一个整型数组，数组里有正数也有负数。数组中一个或连续的多个整
// 数组成一个子数组。求所有子数组的和的最大值。要求时间复杂度为O(n)。

func findMaxSumOfSubArray(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	curSum := arr[0]
	maxSum := arr[0]

	for i := 1; i < len(arr); i++ {
		if curSum < 0 {
			curSum = arr[i]
		} else {
			curSum += arr[i]
		}

		if curSum > maxSum {
			maxSum = curSum
		}
	}

	return maxSum
}
