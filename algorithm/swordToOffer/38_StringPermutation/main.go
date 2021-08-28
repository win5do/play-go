package main

import (
	"fmt"
)

// 面试题38：字符串的排列
// 题目：输入一个字符串，打印出该字符串中字符的所有排列。例如输入字符串abc，
// 则打印出由字符a、b、c所能排列出来的所有字符串abc、acb、bac、bca、cab和cba。

func permutation(str string) {
	if len(str) == 0 {
		return
	}

	rs := []rune(str)
	permutationCore([]rune{}, rs)
}

func permutationCore(begin, rest []rune) {
	if len(rest) == 1 {
		begin = append(begin, rest...)
		fmt.Println(string(begin))
		return
	}

	for i, v := range rest {
		// 要copy slice不然会错乱
		newBegin := make([]rune, len(begin)+1)
		copy(newBegin, begin)
		newRest := make([]rune, len(rest))
		copy(newRest, rest)

		newRest[0], newRest[i] = newRest[i], newRest[0]
		newBegin = append(newBegin, v)
		permutationCore(newBegin, newRest[1:])
	}
}
