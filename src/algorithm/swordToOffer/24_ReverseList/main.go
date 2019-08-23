package main

import (
	"playGo/src/algorithm/swordToOffer/dataStruct/linkedList"
)

// --- 循环实现 ---
func reverseList(pHead *linkedList.ListNode) *linkedList.ListNode {
	var pPrev *linkedList.ListNode
	pNode := pHead

	for pNode != nil {
		pNext := pNode.Next
		pNode.Next = pPrev
		pPrev = pNode
		pNode = pNext
	}

	return pPrev
}

// --- 递归实现 ---
func reverseList_recurse(pHead *linkedList.ListNode) *linkedList.ListNode {
	if pHead == nil || pHead.Next == nil {
		return pHead
	}

	node := reverseList_recurse(pHead.Next)
	pHead.Next.Next = pHead
	pHead.Next = nil
	return node
}
