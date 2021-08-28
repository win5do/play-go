package main

// 面试题10：斐波那契数列
// 题目：写一个函数，输入n，求斐波那契（Fibonacci）数列的第n项。

// --- slow ---
func fibonacciSlow(x int) int {
	if x <= 0 {
		return 0
	}

	if x == 1 {
		return 1
	}

	return fibonacciSlow(x-1) + fibonacciSlow(x-2)
}

// --- loop ---
func fibonacciLoop(x int) int {
	a, b := 1, 0

	for i := 0; i < x; i++ {
		a, b = b, a+b
	}
	return b
}

// --- recurse ---
func fibonacciRecurse(x int) int {
	return fibCore(1, 0, x)
}

func fibCore(a, b, x int) int {
	if x <= 0 {
		return b
	}

	return fibCore(b, a+b, x-1)
}
