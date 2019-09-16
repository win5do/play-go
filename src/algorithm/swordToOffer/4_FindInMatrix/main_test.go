package main

import (
	"fmt"
	"testing"
)

func TestFunc(t *testing.T) {
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
