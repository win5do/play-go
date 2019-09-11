package main

import (
	"playGo/src/algorithm/swordToOffer/dataStruct/tree"
)

// 面试题54：二叉搜索树的第k个结点
// 题目：给定一棵二叉搜索树，请找出其中的第k大的结点。

func kthNodeInBST(pHead *tree.Tree, k int) *tree.Tree {
	if pHead == nil || k <= 0 {
		return nil
	}

	return kthNodeCore(pHead, &k)
}

func kthNodeCore(pNode *tree.Tree, k *int) *tree.Tree {
	if pNode == nil {
		return nil
	}

	var target *tree.Tree

	target = kthNodeCore(pNode.Left, k)
	if target != nil {
		return target
	}

	*k--
	if *k == 0 {
		return pNode
	}

	target = kthNodeCore(pNode.Right, k)

	return target
}
