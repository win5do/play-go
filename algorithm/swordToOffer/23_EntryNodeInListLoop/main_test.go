package main

import (
	"testing"

	"play-go/algorithm/swordToOffer/dataStruct/list"

	"github.com/stretchr/testify/assert"
)

func TestEntryNodeInListLoop(t *testing.T) {

	type casea struct {
		pHead    *list.ListNode
		expected *list.ListNode
	}

	var cs []casea

	addCase := func(a casea) {
		cs = append(cs, a)
	}
	var f func() casea

	// A-A
	f = func() casea {
		p1 := list.CreateListNode(1)
		list.ConnectListNodes(p1, p1)
		return casea{
			pHead:    p1,
			expected: p1,
		}
	}
	addCase(f())

	// A-B-A
	f = func() casea {
		p1 := list.CreateListNode(1)
		p2 := list.CreateListNode(2)
		list.ConnectListNodes(p1, p2)
		list.ConnectListNodes(p2, p1)
		return casea{
			pHead:    p1,
			expected: p1,
		}
	}
	addCase(f())

	// A-B-C 无环
	f = func() casea {
		p1 := list.CreateListNode(1)
		p2 := list.CreateListNode(2)
		p3 := list.CreateListNode(3)
		list.ConnectListNodes(p1, p2)
		list.ConnectListNodes(p2, p3)
		return casea{
			pHead:    p1,
			expected: nil,
		}
	}
	addCase(f())

	// A-B-C
	f = func() casea {
		p1 := list.CreateListNode(1)
		p2 := list.CreateListNode(2)
		p3 := list.CreateListNode(3)
		list.ConnectListNodes(p1, p2)
		list.ConnectListNodes(p2, p3)
		list.ConnectListNodes(p3, p1)
		return casea{
			pHead:    p1,
			expected: p1,
		}
	}
	addCase(f())

	// A-B-C-D-B
	f = func() casea {
		p1 := list.CreateListNode(1)
		p2 := list.CreateListNode(2)
		p3 := list.CreateListNode(3)
		p4 := list.CreateListNode(4)
		list.ConnectListNodes(p1, p2)
		list.ConnectListNodes(p2, p3)
		list.ConnectListNodes(p3, p4)
		list.ConnectListNodes(p4, p2)
		return casea{
			pHead:    p1,
			expected: p2,
		}
	}
	addCase(f())

	for i, v := range cs {
		r := entryNodeInListLoop(v.pHead)
		assert.Equal(t, v.expected, r, i)
	}
}
