package main

import (
	"playGo/src/algorithm/swordToOffer/dataStruct/tree"
)

// 面试题36：二叉搜索树与双向链表
// 题目：输入一棵二叉搜索树，将该二叉搜索树转换成一个排序的双向链表。要求
// 不能创建任何新的结点，只能调整树中结点指针的指向。

func covert(pHead *tree.Tree) *tree.Tree {
	if pHead == nil {
		return nil
	}

	pHead, _ = convertCore(pHead)
	return pHead
}

func convertCore(pNode *tree.Tree) (*tree.Tree, *tree.Tree) {
	pFirst, pLast := pNode, pNode

	if pNode.Right != nil {
		pNode.Right, pLast = convertCore(pNode.Right)
		pNode.Right.Left = pNode
	}

	if pNode.Left != nil {
		pFirst, pNode.Left = convertCore(pNode.Left)
		pNode.Left.Right = pNode
	}

	return pFirst, pLast
}
