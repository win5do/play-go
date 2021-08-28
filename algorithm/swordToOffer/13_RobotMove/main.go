package main

// 面试题13：机器人的运动范围
// 题目：地上有一个m行n列的方格。一个机器人从坐标(0, 0)的格子开始移动，它
// 每一次可以向左、右、上、下移动一格，但不能进入行坐标和列坐标的数位之和
// 大于k的格子。例如，当k为18时，机器人能够进入方格(35, 37)，因为3+5+3+7=18。
// 但它不能进入方格(35, 38)，因为3+5+3+8=19。请问该机器人能够到达多少个格子？

func movingCount(threshold, rows, cols int) int {
	if threshold < 0 || rows <= 0 || cols == 0 {
		return 0
	}

	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	count := movingCountCore(threshold, visited, 0, 0)
	return count
}

func movingCountCore(threshold int, visited [][]bool, row, col int) int {
	count := 0
	if check(threshold, visited, row, col) {
		visited[row][col] = true
		count = 1 +
			movingCountCore(threshold, visited, row-1, col) +
			movingCountCore(threshold, visited, row+1, col) +
			movingCountCore(threshold, visited, row, col-1) +
			movingCountCore(threshold, visited, row, col+1)
	}

	return count
}

func check(threshold int, visited [][]bool, row, col int) bool {
	if row < 0 || col < 0 || row >= len(visited) || col >= len(visited[0]) ||
		visited[row][col] ||
		getDigitSum(row)+getDigitSum(col) > threshold {
		return false
	}
	return true
}

func getDigitSum(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n = n / 10
	}
	return sum
}
