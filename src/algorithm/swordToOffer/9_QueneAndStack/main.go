package main

import (
	"playGo/src/algorithm/swordToOffer/dataStruct/stack"
)

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
