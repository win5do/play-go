package main

// 面试题21：调整数组顺序使奇数位于偶数前面
// 题目：输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有
// 奇数位于数组的前半部分，所有偶数位于数组的后半部分。

func ReorderArray(arr []int) {
	if len(arr) == 0 {
		return
	}

	pBegin := 0
	pEnd := len(arr) - 1

	for pBegin < pEnd {
		if !isOdd(arr[pBegin]) && isOdd(arr[pEnd]) {
			arr[pBegin], arr[pEnd] = arr[pEnd], arr[pBegin]
		}

		if isOdd(arr[pBegin]) {
			pBegin++
		}
		if !isOdd(arr[pEnd]) {
			pEnd--
		}
	}
}

func isOdd(n int) bool {
	return n&1 == 1
}
