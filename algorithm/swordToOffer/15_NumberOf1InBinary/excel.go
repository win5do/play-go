package main

// 面试题：excel中用A表示第1列，B表示第2列......Z表示第26列，AA表示第27列，AB表示28列。请实现一个函数输入字母变化，输出他是第几列。

// 本质上是把十进制数用A~Z表示成26进制。这个26进制是从1开始，而不是0，所以是满27进1。
var alphabetMap = make(map[rune]int, 26)

func init() {
	for i, v := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		alphabetMap[v] = i + 1
	}
}

func excelNumber(str string) int {
	rs := []rune(str)
	result := 0
	base := 1
	bit := len(alphabetMap) // *26

	for i := len(str) - 1; i >= 0; i-- {
		n, ok := alphabetMap[rs[i]]
		if !ok {
			panic("invalid input")
		}

		result += n * base
		base *= bit // 26^n
	}

	return result
}
