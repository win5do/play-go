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

// loop
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head

	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}

	return pre
}

func reverseListRecurse(head *ListNode) *ListNode {
	return recurse(nil, head)
}

func recurse(pre, cur *ListNode) *ListNode {
	if cur == nil {
		return pre
	}

	tmp := cur.Next
	cur.Next = pre
	pre = cur
	cur = tmp

	return recurse(pre, cur)
}
