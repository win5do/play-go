package main

import (
	"fmt"
)

type Stack struct {
	slice []int
}

func (s *Stack) Push(val int) {
	s.slice = append(s.slice, val)
}

func (s *Stack) Pop() (int, bool) {
	l := len(s.slice)
	if l == 0 {
		return -0, false
	}

	val := s.slice[l-1]
	s.slice = s.slice[:l-1]
	return val, true
}

func reverseByStack(pHead *ListNode) {
	stack := new(Stack)

	pNode := pHead
	for pNode != nil {
		stack.Push(pNode.Val)
		pNode = pNode.Next
	}

	ok := true
	for ok {
		var v int
		v, ok = stack.Pop()
		if ok {
			fmt.Println(v)
		}
	}
}

func reverseByRecursion(pHead *ListNode) {
	if pHead == nil {
		return
	}

	if pHead.Next != nil {
		reverseByRecursion(pHead.Next)
	}

	fmt.Println(pHead.Val)
}

func reverseLinkedList(pHead *ListNode) *ListNode {
	if pHead == nil {
		return nil
	}

	var pNodePrev *ListNode
	pNode := pHead

	for pNode != nil {
		pNodeNext := pNode.Next
		pNode.Next = pNodePrev
		pNodePrev = pNode
		pNode = pNodeNext
	}

	return pNodePrev
}
