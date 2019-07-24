package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddToTail(pHead *ListNode, val int) {
	pNew := new(ListNode)
	pNew.Val = val
	pNew.Next = nil

	if pHead == nil {
		pHead = pNew
		return
	}

	pNode := pHead.Next
	for pNode.Next != nil {
		pNode = pNode.Next
	}

	pNode.Next = pNew
}

func RemoveNode(pHead *ListNode, val int) *ListNode {
	if pHead == nil {
		return nil
	}

	if pHead.Val == val {
		pHead = pHead.Next
		// 必须使用返回值，因为go是值传递没办法把head设为nil，除非使用**LinkedNode
		return pHead
	}

	pNode := pHead
	pNodeNext := pNode.Next

	for pNodeNext != nil {
		if pNodeNext.Val == val {
			pNode.Next = pNodeNext.Next
			return pHead
		}

		pNode = pNodeNext
		pNodeNext = pNodeNext.Next
	}

	return pHead
}

func PrintList(pHead *ListNode) {
	pNode := pHead
	for pNode != nil {
		fmt.Println(pNode.Val)
		pNode = pNode.Next
	}
}

func makeList(x int) *ListNode {
	if x < 1 {
		return nil
	}

	pHead := &ListNode{
		Val: 0,
	}

	i := 1
	pNode := pHead
	for i < x {
		pNode.Next = &ListNode{
			Val: i,
		}
		pNode = pNode.Next
		i++
	}

	return pHead
}
