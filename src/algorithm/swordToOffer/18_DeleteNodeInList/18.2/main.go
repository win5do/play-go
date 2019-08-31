package main

import (
	"playGo/src/algorithm/swordToOffer/dataStruct/list"
)

// 面试题18（二）：删除链表中重复的结点
// 题目：在一个排序的链表中，如何删除重复的结点？例如，在图3.4（a）中重复
// 结点被删除之后，链表如图3.4（b）所示。

func deleteDuplication(pHead *list.ListNode) {
	if pHead == nil {
		return
	}

	prevNode := pHead
	pNode := pHead.Next

	for pNode != nil {
		if pNode.Val != prevNode.Val {
			prevNode = pNode
			pNode = pNode.Next
		} else {
			// 检查是不是有连续重复的节点，将prevNode指向下一个不重复的节点
			for pNode.Next != nil && pNode.Next.Val == prevNode.Val {
				pNode = pNode.Next
			}

			prevNode.Next = pNode.Next
			pNode = pNode.Next
		}
	}
}
