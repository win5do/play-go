package main

import (
	"playGo/src/algorithm/swordToOffer/dataStruct/linkedList"
)

// 面试题24：反转链表
// 题目：定义一个函数，输入一个链表的头结点，反转该链表并输出反转后链表的
// 头结点。

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
