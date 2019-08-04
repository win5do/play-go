package main

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
