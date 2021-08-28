package main

import (
	"container/list"
)

// 面试题59（一）：滑动窗口的最大值
// 题目：给定一个数组和滑动窗口的大小，请找出所有滑动窗口里的最大值。例如，
// 如果输入数组{2, 3, 4, 2, 6, 2, 5, 1}及滑动窗口的大小3，那么一共存在6个
// 滑动窗口，它们的最大值分别为{4, 4, 6, 6, 6, 5}，

func maxInWindows(arr []int, size int) []int {
	result := make([]int, 0)
	if len(arr) == 0 || size <= 0 {
		return result
	}

	// 存放数组下标
	quene := list.New()

	for i := 0; i < size; i++ {
		if i >= len(arr) {
			max := arr[quene.Front().Value.(int)]
			result = append(result, max)
			return result
		}

		for quene.Len() > 0 && arr[i] >= arr[quene.Back().Value.(int)] {
			quene.Remove(quene.Back())
		}

		quene.PushBack(i)
	}

	max := arr[quene.Front().Value.(int)]
	result = append(result, max)

	for i := size; i < len(arr); i++ {
		if i-quene.Front().Value.(int) >= size {
			quene.Remove(quene.Front())
		}

		for quene.Len() > 0 && arr[i] >= arr[quene.Back().Value.(int)] {
			quene.Remove(quene.Back())
		}

		quene.PushBack(i)

		max := arr[quene.Front().Value.(int)]
		result = append(result, max)
	}

	return result
}
