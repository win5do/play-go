package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}

	a := head
	b := head.Next

	if b == nil {
		return true
	}

	for b.Next != nil {
		if b.Next.Next == nil {
			b = b.Next
		} else {
			b = b.Next.Next
		}

		a = a.Next
	}

	return check(head, reverse(a.Next))
}

func reverse(head *ListNode) *ListNode {
	var a *ListNode
	b := head

	for b != nil {
		tmp := b.Next
		b.Next = a
		a = b
		b = tmp
	}

	return a
}

func check(a, b *ListNode) bool {
	if a == nil || b == nil {
		return a == b
	}

	for a != nil && b != nil {
		if a.Val != b.Val {
			return false
		}
		a = a.Next
		b = b.Next
	}

	return true
}
