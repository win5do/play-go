package main

// 面试题49：丑数
// 题目：我们把只包含因子2、3和5的数称作丑数（Ugly Number）。求按从小到
// 大的顺序的第1500个丑数。例如6、8都是丑数，但14不是，因为它包含因子7。
// 习惯上我们把1当做第一个丑数。

// --- violent solution ---
// 暴力解法
func isUgly(n int) bool {
	for n%2 == 0 {
		n /= 2
	}

	for n%3 == 0 {
		n /= 3
	}

	for n%5 == 0 {
		n /= 5
	}

	return n == 1
}

func getUglyNumber(index int) int {
	count := 0
	n := 1

	for count < index {
		if isUgly(n) {
			count++
		}
		n++
	}

	return n - 1
}

// --- extra array ---
// 使用额外数组空间换时间
func getUglyNumber2(index int) int {
	if index <= 0 {
		return 0
	}

	arr := make([]int, index)
	arr[0] = 1
	nextIndex := 1

	// 记录2,3,5分别乘以2,3,5时刚好比数组最后一个数大时的index
	i2, i3, i5 := 0, 0, 0

	for nextIndex < index {
		mina := min(arr[i2]*2, arr[i3]*3, arr[i5]*5)
		arr[nextIndex] = mina

		for arr[i2]*2 <= mina {
			i2++
		}

		for arr[i3]*3 <= mina {
			i3++
		}

		for arr[i5]*5 <= mina {
			i5++
		}

		nextIndex++
	}

	return arr[nextIndex-1]
}

func min(n1, n2, n3 int) int {
	var min int

	if n1 < n2 {
		min = n1
	} else {
		min = n2
	}

	if n3 < min {
		min = n3
	}

	return min
}
