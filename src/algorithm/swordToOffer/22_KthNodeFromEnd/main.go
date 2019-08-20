package main

import (
	"playGo/src/algorithm/swordToOffer/dataStruct/linkedList"
)

func findKthToTail(pHead *linkedList.ListNode, k int) *linkedList.ListNode {
	if pHead == nil || k <= 0 {
		return nil
	}

	pAhead, pBehind := pHead, pHead
	n := 1

	for pAhead.Next != nil {
		pAhead = pAhead.Next
		n++
		if n > k {
			// p1 - p2 = k -1
			pBehind = pBehind.Next
		}
	}

	if n < k {
		// 链表长度不足k
		return nil
	}

	return pBehind
}
