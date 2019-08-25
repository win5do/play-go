package main

import (
	"fmt"
)

// 面试题29：顺时针打印矩阵
// 题目：输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字。

func printMatrixClockwisely(matrix [][]int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}

	rows := len(matrix)
	cols := len(matrix[0])

	for i := 0; i*i < rows && i*i < cols; i++ {
		printMatrixCircle(matrix, i)
	}
}

func printMatrixCircle(matrix [][]int, start int) {
	rows := len(matrix)
	cols := len(matrix[0])

	if start >= rows || start >= cols {
		return
	}

	endX := cols - start - 1
	endY := rows - start - 1

	// 从左到右打印一行
	for i := start; i < cols-start; i++ {
		fmt.Println(matrix[start][i])
	}

	// 从上到下打印一行
	if endX >= start {
		for i := start + 1; i <= endY; i++ {
			fmt.Println(matrix[i][endX])
		}
	}

	// 从右到左打印一行
	if endY > start {
		for i := endX - 1; i >= start; i-- {
			fmt.Println(matrix[endY][i])
		}
	}

	// 从下到上打印一行
	if endY-1 > start {
		for i := endY - 1; i >= start+1; i-- {
			fmt.Println(matrix[i][start])
		}
	}
}
