package main

import (
	"math"
)

// 面试题14：剪绳子
// 题目：给你一根长度为n绳子，请把绳子剪成m段（m、n都是整数，n>1并且m≥1）。
// 每段的绳子的长度记为k[0]、k[1]、……、k[m]。k[0]*k[1]*…*k[m]可能的最大乘
// 积是多少？例如当绳子的长度是8时，我们把它剪成长度分别为2、3、3的三段，此
// 时得到最大的乘积18。

// 动态规划
func maxProductAfterCutting_dynamic(n int) int {
	if n < 2 {
		return 0
	}
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}

	products := make([]int, n+1)
	// 123不剪则是最优值，题目规定必须剪一刀，<=3的结果在前面做了判断
	products[0] = 0
	products[1] = 1
	products[2] = 2
	products[3] = 3

	for i := 4; i <= n; i++ {
		max := 0
		for j := 1; j <= i/2; j++ {
			r := products[j] * products[i-j]
			if r > max {
				max = r
			}
		}
		products[i] = max
	}

	max := products[n]
	return max
}

// 贪婪算法
func maxProductAfterCutting_greedy(n int) int {
	if n < 2 {
		return 0
	}
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}

	// 证明：n>=5时 2(n-2)>n 3(n-3)>n 3(n-3)>=2(n-2)
	// 所以尽可能的多剪长度为3的绳子段，剩下4时剪成2*2
	timesOf3 := n / 3
	if n-timesOf3*3 == 1 {
		timesOf3 -= 1
	}
	timesOf2 := (n - timesOf3*3) / 2

	max := math.Pow(3, float64(timesOf3)) * math.Pow(2, float64(timesOf2))
	return int(max)
}
