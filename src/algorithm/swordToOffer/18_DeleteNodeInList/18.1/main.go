package main

import (
	"playGo/src/algorithm/swordToOffer/dataStruct/linkedList"
)

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
