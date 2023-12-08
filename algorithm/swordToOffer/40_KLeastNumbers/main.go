package main

import (
	"container/heap"

	"play-go/algorithm/swordToOffer/dataStruct/array"
	"play-go/algorithm/swordToOffer/dataStruct/heapa"
)

// 面试题40：最小的k个数
// 题目：输入n个整数，找出其中最小的k个数。例如输入4、5、1、6、2、7、3、8
// 这8个数字，则最小的4个数字是1、2、3、4。

// --- partition ---
func getLeastNumber_partition(input []int, k int) []int {
	if len(input) < 0 || k < 0 {
		return nil
	}

	start := 0
	end := len(input) - 1
	index := array.Partition(input, start, k)

	for index != k-1 {
		if index > k-1 {
			end := index - 1
			index = array.Partition(input, start, end)
		} else {
			start = index + 1
			index = array.Partition(input, start, end)
		}
	}

	return input[:index]
}

// --- heap ---
func getLeastNumber_heap(input []int, k int) []int {
	if len(input) < 0 || k < 0 {
		return nil
	}

	h := new(heapa.MaxHeap)
	for i, v := range input {
		if i < k {
			heap.Push(h, v)
		} else if v < h.Top() {
			h.Pop()
			heap.Push(h, v)
		}
	}

	return h.Value()
}
