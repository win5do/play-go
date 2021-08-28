package main

import (
	"fmt"
	"testing"
)

func makeMatrix(row, col int) [][]int {
	matrix := make([][]int, row)
	for i := range matrix {
		row := make([]int, col)
		for j := range row {
			row[j] = j + 1 + col*i
		}
		matrix[i] = row
	}
	return matrix
}

func TestPrintMatrixClockwisely(t *testing.T) {
	cs := []struct {
		matrix [][]int
	}{
		{matrix: makeMatrix(4, 4)},
		{matrix: makeMatrix(1, 1)},
		{matrix: makeMatrix(0, 4)},
		{matrix: makeMatrix(4, 0)},
		{matrix: makeMatrix(1, 4)},
		{matrix: makeMatrix(4, 1)},
		{matrix: makeMatrix(4, 5)},
		{matrix: makeMatrix(5, 3)},
		{matrix: makeMatrix(3, 4)},
	}

	for i, v := range cs {
		fmt.Printf("case: %v\n", i)
		printMatrixClockwisely(v.matrix)
	}
}
