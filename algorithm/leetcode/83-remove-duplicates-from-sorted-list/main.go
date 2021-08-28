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

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	pre := head
	cur := head.Next

	for cur != nil {
		if pre.Val == cur.Val {
			pre.Next = cur.Next
			cur = pre.Next
		} else {
			pre = cur
			cur = cur.Next
		}
	}

	return head
}

func removeDuplicatesRecurse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	head.Next = removeDuplicatesRecurse(head.Next)

	if head.Val == head.Next.Val {
		head.Next = head.Next.Next
	}

	return head
}
