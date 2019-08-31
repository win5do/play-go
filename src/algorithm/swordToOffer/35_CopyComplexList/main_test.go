package main

import (
	"fmt"
	"testing"
)

func newNode(n int) *complexListNode {
	return &complexListNode{
		Val: n,
	}
}

func printList(pHead *complexListNode) {
	pNode := pHead
	for pNode != nil {
		sib := -1
		if pNode.Sibling != nil {
			sib = pNode.Sibling.Val
		}
		fmt.Println(pNode.Val, sib)
		pNode = pNode.Next
	}
	fmt.Println("---")
}

func connectNodes(pNode, pNext, pSib *complexListNode) {
	pNode.Next = pNext
	pNode.Sibling = pSib
}

func TestClone(t *testing.T) {
	// nil
	r := clone(nil)
	printList(r)

	// 单节点
	p1 := newNode(1)
	connectNodes(p1, nil, p1)
	r = clone(p1)
	printList(r)

	// 多节点
	p1 = newNode(1)
	p2 := newNode(2)
	p3 := newNode(3)
	p4 := newNode(4)
	p5 := newNode(5)

	connectNodes(p1, p2, p3)
	connectNodes(p2, p3, p5)
	connectNodes(p3, p4, nil)
	connectNodes(p4, p5, p1)

	r = clone(p1)
	printList(r)
	printList(p1)
}
