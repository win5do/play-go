package main

// 面试题12：矩阵中的路径
// 题目：请设计一个函数，用来判断在一个矩阵中是否存在一条包含某字符串所有
// 字符的路径。路径可以从矩阵中任意一格开始，每一步可以在矩阵中向左、右、
// 上、下移动一格。如果一条路径经过了矩阵的某一格，那么该路径不能再次进入
// 该格子。例如在下面的3×4的矩阵中包含一条字符串“bfce”的路径（路径中的字
// 母用下划线标出）。但矩阵中不包含字符串“abfb”的路径，因为字符串的第一个
// 字符b占据了矩阵中的第一行第二个格子之后，路径不能再次进入这个格子。
// A B T G
// C F C S
// J D E H

func hasPath(matrix [][]string, str string) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 || len(str) == 0 {
		return false
	}

	rows := len(matrix)
	cols := len(matrix[0])

	visited := make([][]bool, rows)
	for row := range visited {
		val := make([]bool, cols)
		visited[row] = val
	}

	for row := range matrix {
		for col := range matrix[row] {
			if hasPathCore(matrix, visited, str, 0, row, col) {
				return true
			}
		}
	}

	return false
}

func hasPathCore(matrix [][]string, visited [][]bool, str string, pathLength, row, col int) bool {
	if pathLength == len(str) {
		return true
	}

	hasPath := false
	if row >= 0 && row < len(matrix) && col >= 0 && col < len(matrix[0]) && matrix[row][col] == string(str[pathLength]) && visited[row][col] == false {
		pathLength++
		visited[row][col] = true

		hasPath = hasPathCore(matrix, visited, str, pathLength, row-1, col) ||
			hasPathCore(matrix, visited, str, pathLength, row, col-1) ||
			hasPathCore(matrix, visited, str, pathLength, row+1, col) ||
			hasPathCore(matrix, visited, str, pathLength, row, col+1)

		if !hasPath {
			pathLength--
			visited[row][col] = false
		}
	}

	return hasPath
}
