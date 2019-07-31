package main

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
	if row >= 0 && row < len(matrix) && col >=0 && col < len(matrix[0]) && matrix[row][col] == string(str[pathLength]) && visited[row][col] == false {
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
