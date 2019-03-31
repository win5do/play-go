package linkedList

import (
	"testing"
)

func newList(s []int) *listNode {
	var head *listNode
	pre := new(listNode)

	for k, i := range s {
		cur := new(listNode)
		cur.Val = i
		pre.Next = cur
		pre = cur

		if k == 0 {
			head = cur
		}
	}

	return head
}

func TestNewList(t *testing.T) {
	head := newList([]int{1, 2, 3})
	if head.Val != 1 || head.Next.Val != 2 || head.Next.Next.Val != 3 {
		t.Fatal()
	}
}

func TestReverseList(t *testing.T) {
	head := newList([]int{1, 2, 3})
	newHead := reverseList(head)
	if newHead.Val != 3 || newHead.Next.Val != 2 || newHead.Next.Next.Val != 1 {
		t.Fatal()
	}

	head = newList([]int{1, 2, 3})
	newHead = reverseList2(head)
	if newHead.Val != 3 || newHead.Next.Val != 2 || newHead.Next.Next.Val != 1 {
		t.Fatal()
	}
}

func TestRemoveDuplicates(t *testing.T) {
	head := newList([]int{1, 1, 1})
	head = removeDuplicates(head)
	if head.Val != 1 || head.Next != nil {
		t.Fatal()
	}
}

func TestRemoveDuplicates2(t *testing.T) {
	head := newList([]int{1, 1, 1})
	head = removeDuplicates2(head)
	if head.Val != 1 || head.Next != nil {
		t.Fatal()
	}
}
