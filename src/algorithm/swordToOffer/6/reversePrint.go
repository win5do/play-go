package main

import (
	"fmt"

	"playGo/src/algorithm/swordToOffer/dataStruct/stack"
)

func reverseByStack(pHead *ListNode) {
	stack := new(stack.Stack)

	pNode := pHead
	for pNode != nil {
		stack.Push(pNode.Val)
		pNode = pNode.Next
	}

	for stack.Size() > 0 {
		v := stack.Pop()
		fmt.Println(v)
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
