package main

import (
	"fmt"

	"playGo/src/algorithm/swordToOffer/dataStruct/stack"
)

// 从尾到头打印链表
func reversePrint_stack(pHead *ListNode) {
	stacka := new(stack.Stack)

	pNode := pHead
	for pNode != nil {
		stacka.Push(pNode.Val)
		pNode = pNode.Next
	}

	for stacka.Size() > 0 {
		v := stacka.Pop()
		fmt.Println(v)
	}
}

func reversePrint_recursion(pHead *ListNode) {
	if pHead == nil {
		return
	}

	if pHead.Next != nil {
		reversePrint_recursion(pHead.Next)
	}

	fmt.Println(pHead.Val)
}

// 翻转链表
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
