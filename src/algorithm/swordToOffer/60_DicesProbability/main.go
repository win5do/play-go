package main

import (
	"fmt"
	"math"
)

// 面试题60：n个骰子的点数
// 题目：把n个骰子扔在地上，所有骰子朝上一面的点数之和为s。输入n，打印出s
// 的所有可能的值出现的概率。
const NUMBER = 6

func printProbability_recurse(n int) {
	if n < 1 {
		return
	}

	max := n * 6
	result := make([]int, max-NUMBER+1)

	probability(n, n, 0, result)

	total := math.Pow(NUMBER, float64(n))

	for i, v := range result {
		fmt.Printf("%d: %f\n", i+n, float64(v)/total)
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
