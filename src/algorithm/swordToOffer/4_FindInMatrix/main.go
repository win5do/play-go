package main

import (
	"fmt"
)

func find(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	row := 0
	column := len(matrix[0]) - 1

	for row < len(matrix) && column >= 0 {
		number := matrix[row][column]
		if number == target {
			return true
		} else if number > target {
			column--
		} else {
			row++
		}
	}
	return false
}

func main() {
	matrix := [][]int{
		{1, 2, 3, 4},
		{2, 3, 4, 5},
		{3, 4, 5, 6},
		{4, 5, 6, 7},
	}

	fmt.Println(find(matrix, 7))
	fmt.Println(find(matrix, 1))
	fmt.Println(find(matrix, 3))
	fmt.Println(find(matrix, 0))
	fmt.Println(find(matrix, 10))
}
