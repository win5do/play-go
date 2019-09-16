package main

import (
	"playGo/src/algorithm/swordToOffer/dataStruct/stack"
)

// 面试题9：用两个栈实现队列
// 题目：用两个栈实现一个队列。队列的声明如下，请实现它的两个函数appendTail
// 和deleteHead，分别完成在队列尾部插入结点和在队列头部删除结点的功能。

type Queue struct {
	stackA stack.Stack
	stackB stack.Stack
}

func (q *Queue) Push(x int) {
	q.stackA.Push(x)
}

func (q *Queue) Pop() int {
	if q.stackB.Size() <= 0 {
		for q.stackA.Size() > 0 {
			x := q.stackA.Pop()
			q.stackB.Push(x)
		}
	}

	if q.stackB.Size() <= 0 {
		panic("queue empty")
	}

	return q.stackB.Pop()
}
