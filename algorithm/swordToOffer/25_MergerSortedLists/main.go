package main

import (
	"play-go/algorithm/swordToOffer/dataStruct/list"
)

// 面试题25：合并两个排序的链表
// 题目：输入两个递增排序的链表，合并这两个链表并使新链表中的结点仍然是按
// 照递增排序的。例如输入图3.11中的链表1和链表2，则合并之后的升序链表如链
// 表3所示。

func mergeList(p1, p2 *list.ListNode) *list.ListNode {
	if p1 == nil {
		return p2
	}
	if p2 == nil {
		return p1
	}

	var mergeHead *list.ListNode

	if p2.Val < p1.Val {
		mergeHead = p2
		mergeHead.Next = mergeList(p1, p2.Next)
	} else {
		mergeHead = p1
		mergeHead.Next = mergeList(p1.Next, p2)
	}

	return mergeHead
}
