package main

// 面试题11：旋转数组的最小数字
// 题目：把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。
// 输入一个递增排序的数组的一个旋转，输出旋转数组的最小元素。例如数组
// {3, 4, 5, 1, 2}为{1, 2, 3, 4, 5}的一个旋转，该数组的最小值为1。

func min(arr []int) int {
	if len(arr) == 0 {
		panic("invalid input")
	}

	index1 := 0
	index2 := len(arr) - 1
	indexMid := 0

	// 不能用for index1 < index2作为条件，可能是1，2，3这种顺序数组（旋转了0个元素）
	for arr[index1] >= arr[index2] {
		if index2-index1 == 1 {
			return arr[index2]
		}

		// 如果下标为index1、index2和indexMid指向的三个数字相等，
		// 则只能顺序查找
		if arr[index1] == arr[indexMid] && arr[index2] == arr[indexMid] {
			return minInOrder(arr, index1, index2)
		}

		indexMid = (index1 + index2) >> 1
		if arr[indexMid] >= arr[index1] {
			index1 = indexMid
		} else if arr[indexMid] <= arr[index2] {
			index2 = indexMid
		}
	}

	return arr[indexMid]
}

func minInOrder(arr []int, index1, index2 int) int {
	min := arr[index1]

	for i := index1; i <= index2; i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}

	return min
}
