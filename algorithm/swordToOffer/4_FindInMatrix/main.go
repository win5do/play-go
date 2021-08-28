package main

// 面试题4：二维数组中的查找
// 题目：在一个二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按
// 照从上到下递增的顺序排序。请完成一个函数，输入这样的一个二维数组和一个
// 整数，判断数组中是否含有该整数。

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
