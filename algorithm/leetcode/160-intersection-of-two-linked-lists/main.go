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

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	l := headA
	r := headB

	for l != r {
		if l == nil {
			l = headB
		} else {
			l = l.Next
		}

		if r == nil {
			r = headA
		} else {
			r = r.Next
		}
	}

	return r
}
