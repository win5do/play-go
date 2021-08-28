package main

import (
	"fmt"
	"math"
)

// 面试题60：n个骰子的点数
// 题目：把n个骰子扔在地上，所有骰子朝上一面的点数之和为s。输入n，打印出s
// 的所有可能的值出现的概率。
const NUMBER = 6

// --- recurse ---
func printProbability_recurse(n int) {
	if n < 1 {
		return
	}

	max := n * 6
	result := make([]int, max-NUMBER+1)

	probability(n, n, 0, result)

	total := math.Pow(NUMBER, float64(n))

	for i, v := range result {
		ratio := float64(v) / total
		fmt.Printf("%d: %f\n", i+n, ratio)
	}
}

func probability(original, rest, sum int, result []int) {
	if rest < 1 {
		result[sum-original]++
		return
	}

	for i := 1; i <= NUMBER; i++ {
		probability(original, rest-1, sum+i, result)
	}
}

// --- loop ---
func printProbability_loop(n int) {
	if n < 1 {
		return
	}

	results := make([][]int, 2)
	results[0] = make([]int, n*NUMBER+1)
	results[1] = make([]int, n*NUMBER+1)

	flag := 0

	for i := 1; i <= NUMBER; i++ {
		results[flag][i] = 1
	}

	for i := 2; i <= n; i++ {
		// 0到n-1不可能出现，置空
		for j := 0; j < i; j++ {
			results[1-flag][j] = 0
		}

		for j := i; j <= i*NUMBER; j++ {
			results[1-flag][j] = 0
			// 第j项等于j-1,j-2...j-6的和
			for k := 1; k <= j && k <= NUMBER; k++ {
				results[1-flag][j] += results[flag][j-k]
			}
		}

		// 切换数组
		flag = 1 - flag
	}

	total := math.Pow(NUMBER, float64(n))

	for i := n; i <= n*NUMBER; i++ {
		ratio := float64(results[flag][i]) / total
		fmt.Printf("%d: %f\n", i, ratio)
	}
}
