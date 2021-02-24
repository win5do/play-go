package main

import (
	"fmt"

	"playGo/src/algorithm/swordToOffer/dataStruct/stack"
)

// 面试题6：从尾到头打印链表
// 题目：输入一个链表的头结点，从尾到头反过来打印出每个结点的值。

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

	reversePrint_recursion(pHead.Next)

	fmt.Println(pHead.Val)
}

// 翻转链表
func reverseList(pHead *ListNode) *ListNode {
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
