package main

import (
	"playGo/src/algorithm/swordToOffer/dataStruct/linkedList"
)

// 面试题18（一）：在O(1)时间删除链表结点
// 题目：给定单向链表的头指针和一个结点指针，定义一个函数在O(1)时间删除该
// 结点。

func deleteNode(pHead, toBeDeleted *linkedList.ListNode) {
	if pHead == nil || toBeDeleted == nil {
		return
	}

	if toBeDeleted.Next != nil {
		// 待删除节点不是尾节点，将next复制到toBeDeleted
		next := toBeDeleted.Next
		toBeDeleted.Next = nil
		toBeDeleted.Val = next.Val
		toBeDeleted.Next = next.Next
	} else if toBeDeleted == pHead {
		// 待删除节点是head
		pHead = nil
		toBeDeleted = nil
	} else {
		// 待删除节点是尾节点，但不是头节点，需要遍历获取其父节点
		pNode := pHead
		for pNode.Next != toBeDeleted {
			pNode = pNode.Next
		}

		pNode.Next = nil
		toBeDeleted = nil
	}
}
