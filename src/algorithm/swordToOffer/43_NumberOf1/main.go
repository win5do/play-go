package main

// 面试题43：从1到n整数中1出现的次数
// 题目：输入一个整数n，求从1到n这n个整数的十进制表示中1出现的次数。例如
// 输入12，从1到12这些整数中包含1 的数字有1，10，11和12，1一共出现了5次。

// --- method-1 ---
func numberOf1(n uint64) int {
	count := 0
	for n > 0 {
		if n%10 == 1 {
			count++
		}
		n = n / 10
	}
	return count
}

func numberOf1ToN(n uint64) int {
	count := 0

	for i := uint64(1); i <= n; i++ {
		count += numberOf1(i)
	}
	return count
}

// --- method-2 ---
// TODO
