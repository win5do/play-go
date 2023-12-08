package main

import (
	"play-go/algorithm/swordToOffer/dataStruct/stack"
)

// 面试题30：包含min函数的栈
// 题目：定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的min
// 函数。在该栈中，调用min、push及pop的时间复杂度都是O(1)。

type stackWithMin struct {
	data stack.Stack
	min  stack.Stack
}

func (s *stackWithMin) Pop() int {
	s.min.Pop()
	return s.data.Pop()
}

func (s *stackWithMin) Push(val int) {
	s.data.Push(val)

	// 为空时不需要判断大小
	if s.min.Size() == 0 || val < s.min.Top() {
		s.min.Push(val)
	} else {
		s.min.Push(s.min.Top())
	}
}

func (s *stackWithMin) Size() int {
	return s.data.Size()
}

func (s *stackWithMin) Top() int {
	return s.data.Top()
}

func (s *stackWithMin) Min() int {
	return s.min.Top()
}
