package main

import (
	"playGo/src/algorithm/swordToOffer/dataStruct/list"
)

// 面试题52：两个链表的第一个公共结点
// 题目：输入两个链表，找出它们的第一个公共结点。

func findFirstCommonNode(pHead1, pHead2 *list.ListNode) *list.ListNode {
	if pHead1 == nil || pHead2 == nil {
		return nil
	}

	length1 := getListLength(pHead1)
	length2 := getListLength(pHead2)

	var pLong, pShort *list.ListNode

	diff := length1 - length2

	if diff >= 0 {
		pLong = pHead1
		pShort = pHead2
	} else {
		pLong = pHead2
		pShort = pHead1
		diff = -diff
	}

	for i := 0; i < diff; i++ {
		pLong = pLong.Next
	}

	for pLong != pShort {
		pLong = pLong.Next
		pShort = pShort.Next
	}

	return pLong
}

func getListLength(pHead *list.ListNode) int {
	length := 0

	pNode := pHead

	for pNode != nil {
		length++
		pNode = pNode.Next
	}

	return length
}
