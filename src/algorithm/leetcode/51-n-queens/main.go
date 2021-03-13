package main

func solveNQueens(n int) [][]string {
	matrix := make([]int, n)
	var ans [][]string
	calNQueens(0, matrix, &ans)
	return ans
}

func calNQueens(row int, matrix []int, ans *[][]string) {
	n := len(matrix)

	if row == n {
		*ans = append(*ans, toString(matrix))
		return
	}

	for column := 0; column < n; column++ {
		if isOK(matrix, row, column) {
			matrix[row] = column
			calNQueens(row+1, matrix, ans)
		}
	}
}

func isOK(matrix []int, row, column int) bool {
	for i := row - 1; i >= 0; i-- {
		if matrix[i] == column ||
			row-i == abs(column-matrix[i]) {
			return false
		}
	}

	return true
}

func abs(in int) int {
	if in < 0 {
		return -in
	}

	return in
}

func toString(in []int) []string {
	leng := len(in)
	r := make([]string, leng)

	emptyStr := make([]byte, leng)
	for i := range emptyStr {
		emptyStr[i] = '.'
	}

	for i, v := range in {
		str := make([]byte, leng)
		copy(str, emptyStr)
		str[v] = 'Q'
		r[i] = string(str)
	}

	return r
}
