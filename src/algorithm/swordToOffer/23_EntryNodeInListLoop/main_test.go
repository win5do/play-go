package main

import (
	"testing"

	"playGo/src/algorithm/swordToOffer/dataStruct/linkedList"

	"github.com/stretchr/testify/assert"
)

func TestEntryNodeInListLoop(t *testing.T) {

	type casea struct {
		pHead    *linkedList.ListNode
		expected *linkedList.ListNode
	}

	var cs []casea

	addCase := func(a casea) {
		cs = append(cs, a)
	}
	var f func() casea

	// A-A
	f = func() casea {
		p1 := linkedList.CreateListNode(1)
		linkedList.ConnectListNodes(p1, p1)
		return casea{
			pHead:    p1,
			expected: p1,
		}
	}
	addCase(f())

	// A-B-A
	f = func() casea {
		p1 := linkedList.CreateListNode(1)
		p2 := linkedList.CreateListNode(2)
		linkedList.ConnectListNodes(p1, p2)
		linkedList.ConnectListNodes(p2, p1)
		return casea{
			pHead:    p1,
			expected: p1,
		}
	}
	addCase(f())

	// A-B-C 无环
	f = func() casea {
		p1 := linkedList.CreateListNode(1)
		p2 := linkedList.CreateListNode(2)
		p3 := linkedList.CreateListNode(3)
		linkedList.ConnectListNodes(p1, p2)
		linkedList.ConnectListNodes(p2, p3)
		return casea{
			pHead:    p1,
			expected: nil,
		}
	}
	addCase(f())

	// A-B-C
	f = func() casea {
		p1 := linkedList.CreateListNode(1)
		p2 := linkedList.CreateListNode(2)
		p3 := linkedList.CreateListNode(3)
		linkedList.ConnectListNodes(p1, p2)
		linkedList.ConnectListNodes(p2, p3)
		linkedList.ConnectListNodes(p3, p1)
		return casea{
			pHead:    p1,
			expected: p1,
		}
	}
	addCase(f())

	// A-B-C-D-B
	f = func() casea {
		p1 := linkedList.CreateListNode(1)
		p2 := linkedList.CreateListNode(2)
		p3 := linkedList.CreateListNode(3)
		p4 := linkedList.CreateListNode(4)
		linkedList.ConnectListNodes(p1, p2)
		linkedList.ConnectListNodes(p2, p3)
		linkedList.ConnectListNodes(p3, p4)
		linkedList.ConnectListNodes(p4, p2)
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
