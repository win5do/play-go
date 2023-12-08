package main

import (
	"play-go/algorithm/swordToOffer/dataStruct/list"
)

// --- 使用两个指针遍历一次 ---
func findKthToTail(pHead *list.ListNode, k int) *list.ListNode {
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

// --- 递归实现 ---
// 同理还可以用栈等额外空间，没有上面的方法好
func findKthToTail_recurse(pHead *list.ListNode, k int) *list.ListNode {
	if pHead == nil || k <= 0 {
		return nil
	}

	result, _ := traversal(pHead, k)

	return result
}

func traversal(pHead *list.ListNode, k int) (*list.ListNode, int) {
	if pHead == nil {
		return nil, 0
	}

	result, n := traversal(pHead.Next, k)
	n++

	if n == k {
		result = pHead
	}

	return result, n
}
