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

func deleteNode(node *ListNode) {
	if node == nil || node.Next == nil {
		return
	}

	next := node.Next

	node.Val = next.Val
	node.Next = next.Next
}
