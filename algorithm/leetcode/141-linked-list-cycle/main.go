package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	l := head
	r := head.Next
	for l != nil && r != nil {
		if l == r {
			return true
		}

		l = l.Next
		r = r.Next
		if r == nil {
			return false
		}
		r = r.Next
	}

	return false
}
