package main

import (
	"fmt"
	"testing"
)

func TestAddToTail(t *testing.T) {
	pHead := &LinkedNode{
		Val: 0,
	}
	pHead.Next = &LinkedNode{
		Val: 1,
	}

	cs := []*LinkedNode{
		pHead,
		nil,
	}

	for _, v := range cs {
		AddToTail(v, 2)
		PrintListNode(v)
	}
}

func TestRemoveNode(t *testing.T) {
	pHead := &LinkedNode{
		Val: 0,
	}
	pHead.Next = &LinkedNode{
		Val: 1,
	}
	pHead.Next.Next = &LinkedNode{
		Val: 2,
	}

	cs := []*LinkedNode{
		pHead,
		nil,
	}

	for _, v := range cs {
		v = RemoveNode(v, 0)
		PrintListNode(v)
		fmt.Println("-----")
		v = RemoveNode(v, 2)
		PrintListNode(v)
		fmt.Println("-----")
		v = RemoveNode(v, 3)
		PrintListNode(v)
	}
}
