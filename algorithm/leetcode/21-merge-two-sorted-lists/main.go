package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	r := &ListNode{}
	cur := r

	for {
		if l1 == nil {
			cur.Next = l2
			break
		}

		if l2 == nil {
			cur.Next = l1
			break
		}

		var next *ListNode
		if l1.Val < l2.Val {
			next = l1
			l1 = l1.Next
		} else {
			next = l2
			l2 = l2.Next
		}

		cur.Next = next
		cur = next
	}

	return r.Next
}
