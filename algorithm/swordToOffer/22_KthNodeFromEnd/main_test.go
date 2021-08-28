package main

import (
	"testing"

	"playGo/algorithm/swordToOffer/dataStruct/list"

	"github.com/stretchr/testify/assert"
)

func TestFindKthToTail(t *testing.T) {
	p1 := list.CreateListNode(1)
	p2 := list.CreateListNode(2)
	p3 := list.CreateListNode(3)
	p4 := list.CreateListNode(4)
	p5 := list.CreateListNode(5)

	list.ConnectListNodes(p1, p2)
	list.ConnectListNodes(p2, p3)
	list.ConnectListNodes(p3, p4)
	list.ConnectListNodes(p4, p5)

	cs := []struct {
		pHead    *list.ListNode
		k        int
		expected *list.ListNode
	}{
		{pHead: p1, k: 1, expected: p5},
		{pHead: p1, k: 5, expected: p1},
		{pHead: p1, k: 3, expected: p3},
		{pHead: p1, k: 0, expected: nil},
		{pHead: p1, k: 10, expected: nil},
		{pHead: nil, k: 1, expected: nil},
	}

	for i, v := range cs {
		r1 := findKthToTail(v.pHead, v.k)
		r2 := findKthToTail_recurse(v.pHead, v.k)
		assert.Equal(t, v.expected, r1, i)
		assert.Equal(t, r1, r2, i)
	}
}
