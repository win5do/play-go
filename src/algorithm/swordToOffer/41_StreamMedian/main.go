package main

import (
	"container/heap"

	"playGo/src/algorithm/swordToOffer/dataStruct/heapa"
)

// 面试题41：数据流中的中位数
// 题目：如何得到一个数据流中的中位数？如果从数据流中读出奇数个数值，那么
// 中位数就是所有数值排序之后位于中间的数值。如果从数据流中读出偶数个数值，
// 那么中位数就是所有数值排序之后中间两个数的平均值。

type streamReader struct {
	count   int
	minHeap *heapa.MinHeap
	maxHeap *heapa.MaxHeap
}

func NewStreamReader() *streamReader {
	return &streamReader{
		minHeap: new(heapa.MinHeap),
		maxHeap: new(heapa.MaxHeap),
	}
}

func (s *streamReader) insert(n int) {
	s.count++
	if s.count&1 == 1 {
		// 奇数
		if s.minHeap.Len() > 0 && n > s.minHeap.Top() {
			x := heap.Pop(s.minHeap)
			heap.Push(s.maxHeap, x)
			heap.Push(s.minHeap, n)
		} else {
			heap.Push(s.maxHeap, n)
		}
	} else {
		// 偶数
		if s.maxHeap.Len() > 0 && n < s.maxHeap.Top() {
			x := heap.Pop(s.maxHeap)
			heap.Push(s.minHeap, x)
			heap.Push(s.maxHeap, n)
		} else {
			heap.Push(s.minHeap, n)
		}
	}
}

func (s *streamReader) getMedian() float64 {
	if s.count == 0 {
		return 0
	} else if s.count&1 == 1 {
		r := s.maxHeap.Top()
		return float64(r)
	} else {
		r := float64(s.maxHeap.Top()+s.minHeap.Top()) / 2
		return r
	}
}
