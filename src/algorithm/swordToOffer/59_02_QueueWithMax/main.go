package main

import (
	"container/list"
)

// 面试题59（二）：队列的最大值
// 题目：请定义一个队列并实现函数max得到队列里的最大值，
// 要求函数max、push_back和pop_front的时间复杂度都是O(1)。

type queneWithMax struct {
	// 维护两个队列，一个放原始数据，一个放最大值
	data         list.List
	max          list.List
	currentIndex int
}

type element struct {
	value int
	index int
}

func (q *queneWithMax) Len() int {
	return q.data.Len()
}

func (q *queneWithMax) Max() int {
	return q.max.Front().Value.(element).value
}

func (q *queneWithMax) PushBack(e element) {
	for q.max.Len() > 0 && q.max.Front().Value.(element).value >= e.value {
		q.max.Remove(q.max.Front())
	}

	e.index = q.currentIndex
	q.max.PushBack(e)
	q.data.PushBack(e)
	q.currentIndex++
}

func (q *queneWithMax) PopFront() element {
	e := q.data.Remove(q.data.Front()).(element)
	if e.index == q.max.Front().Value.(element).index {
		q.max.Remove(q.max.Front())
	}

	return e
}
