package main

// 面试题35：复杂链表的复制
// 题目：请实现函数ComplexListNode* Clone(ComplexListNode* pHead)，复
// 制一个复杂链表。在复杂链表中，每个结点除了有一个m_pNext指针指向下一个
// 结点外，还有一个m_pSibling 指向链表中的任意结点或者nullptr。

type complexListNode struct {
	Val     int
	Next    *complexListNode
	Sibling *complexListNode
}

func clone(pHead *complexListNode) *complexListNode {
	if pHead == nil {
		return nil
	}

	copyNodes(pHead)
	connectSibling(pHead)
	return divideList(pHead)
}

// clone每一个节点，使node->node'
func copyNodes(pHead *complexListNode) {
	pNode := pHead
	for pNode != nil {
		pCopy := new(complexListNode)
		pCopy.Val = pNode.Val
		pCopy.Next = pNode.Next
		pCopy.Sibling = pNode.Sibling
		pNode.Next = pCopy
		pNode = pCopy.Next
	}
}

// 对node'，将其sibling->sibling'
func connectSibling(pHead *complexListNode) {
	pNode := pHead

	for pNode != nil && pNode.Next != nil {
		if pNode.Sibling != nil {
			pNode.Sibling = pNode.Sibling.Next
		}
		pNode = pNode.Next.Next
	}
}

// 将node和node'分开
func divideList(pHead *complexListNode) *complexListNode {
	pClone := pHead.Next

	pNode := pHead
	for pNode != nil {
		if pNode.Next != nil {
			pNode.Next = pNode.Next.Next
		}
		pNode = pNode.Next
	}

	return pClone
}
