package linkedList

func reverseList_recurse(head *listNode) *listNode {
	if head == nil || head.Next == nil {
		return head
	}

	node := reverseList_recurse(head.Next)
	head.Next.Next = head
	head.Next = nil
	return node
}

func reverseList_loop(head *listNode) *listNode {
	var pre *listNode
	for head != nil {
		next := head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}
