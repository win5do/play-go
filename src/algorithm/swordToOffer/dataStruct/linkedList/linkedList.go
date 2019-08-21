package linkedList

type ListNode struct {
	Val  int
	Next *ListNode
}

func CreateListNode(n int) *ListNode {
	return &ListNode{
		Val: n,
	}
}

func ConnectListNodes(pNode, pNext *ListNode) {
	if pNode == nil {
		return
	}
	pNode.Next = pNext
}
