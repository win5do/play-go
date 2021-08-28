package main

import (
	"playGo/algorithm/swordToOffer/dataStruct/list"
)

func meetingNode(pHead *list.ListNode) (*list.ListNode, int) {
	if pHead == nil {
		return nil, 0
	}
	pAhead, pBehind := pHead, pHead

	var meetingNode *list.ListNode
	for pAhead.Next != nil && pAhead.Next.Next != nil {
		pAhead = pAhead.Next.Next
		pBehind = pBehind.Next

		if pAhead == pBehind {
			meetingNode = pAhead
			break
		}
	}

	if meetingNode == nil {
		return nil, 0
	}

	n := 0
	for pBehind.Next != nil {
		pBehind = pBehind.Next
		n++
		if pBehind == meetingNode {
			break
		}
	}

	return meetingNode, n
}

func entryNodeInListLoop(pHead *list.ListNode) *list.ListNode {
	if pHead == nil {
		return nil
	}

	_, n := meetingNode(pHead)
	// n可能为1，即pHead->Next=pHead
	if n == 0 {
		return nil
	}

	pAhead, pBehind := pHead, pHead

	// pAhead走完环会回到入口需要多走一个环的距离
	for i := 0; i < n; i++ {
		pAhead = pAhead.Next
	}

	// pAhead，pBehind一起走0->n一定会在环的入口相遇
	for pAhead != pBehind {
		pAhead = pAhead.Next
		pBehind = pBehind.Next
	}
	return pAhead
}
