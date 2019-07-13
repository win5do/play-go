package main

import (
	"fmt"
	"testing"
)

func TestAddToTail(t *testing.T) {
	pHead := makeList(10)

	cs := []*ListNode{
		pHead,
		nil,
	}

	for _, v := range cs {
		AddToTail(v, 10)
		PrintList(v)
	}
}

func TestRemoveNode(t *testing.T) {
	pHead := makeList(3)

	cs := []*ListNode{
		pHead,
		nil,
	}

	for _, v := range cs {
		v = RemoveNode(v, 0)
		PrintList(v)
		fmt.Println("-----")
		v = RemoveNode(v, 2)
		PrintList(v)
		fmt.Println("-----")
		v = RemoveNode(v, 3)
		PrintList(v)
	}
}
