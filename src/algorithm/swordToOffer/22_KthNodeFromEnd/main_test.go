package main

import (
	"testing"

	"playGo/src/algorithm/swordToOffer/dataStruct/linkedList"

	"github.com/stretchr/testify/assert"
)

func TestFindKthToTail(t *testing.T) {
	p1 := linkedList.CreateListNode(1)
	p2 := linkedList.CreateListNode(2)
	p3 := linkedList.CreateListNode(3)
	p4 := linkedList.CreateListNode(4)
	p5 := linkedList.CreateListNode(5)

	linkedList.ConnectListNodes(p1, p2)
	linkedList.ConnectListNodes(p2, p3)
	linkedList.ConnectListNodes(p3, p4)
	linkedList.ConnectListNodes(p4, p5)

	cs := []struct {
		pHead    *linkedList.ListNode
		k        int
		expected *linkedList.ListNode
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
