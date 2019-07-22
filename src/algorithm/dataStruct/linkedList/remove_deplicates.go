package linkedList

func removeDuplicates_recurse(head *listNode) *listNode {
	if head == nil || head.Next == nil {
		return head
	}
	head.Next = removeDuplicates_recurse(head.Next)

	if head.Val == head.Next.Val {
		head.Next = head.Next.Next
	}

	return head
}

func removeDuplicates_loop(head *listNode) *listNode {
	cur := head
	for cur != nil && cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}
