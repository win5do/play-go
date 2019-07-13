package main

import (
	"fmt"
)

type LinkedNode struct {
	Val  int
	Next *LinkedNode
}

func AddToTail(pHead *LinkedNode, val int) {
	pNew := new(LinkedNode)
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

func RemoveNode(pHead *LinkedNode, val int) *LinkedNode {
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

func PrintListNode(pHead *LinkedNode) {
	pNode := pHead
	for pNode != nil {
		fmt.Println(pNode.Val)
		pNode = pNode.Next
	}
}
