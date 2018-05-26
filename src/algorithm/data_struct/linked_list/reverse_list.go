package linked_list

func reverseList(head *listNode) *listNode {
	if head.Next == nil {
		return head
	}

	node := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return node
}

func reverseList2(head *listNode) *listNode {
	var pre *listNode
	for head != nil {
		next := head.Next;
		head.Next = pre;
		pre = head;
		head = next;
	}
	return pre;
}
