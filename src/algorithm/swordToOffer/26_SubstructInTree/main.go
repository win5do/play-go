package main

import (
	"playGo/src/algorithm/swordToOffer/dataStruct/tree"
)

// 面试题26：树的子结构
// 题目：输入两棵二叉树A和B，判断B是不是A的子结构。

func hasSubtree(pHead, pSub *tree.BinaryTreeNode) bool {
	if pHead == nil || pSub == nil {
		return false
	}

	if pHead.Val == pSub.Val {
		result := doesTree1HasTree2(pHead, pSub)
		if result {
			return result
		}
	}

	return hasSubtree(pHead.Left, pSub) || hasSubtree(pHead.Right, pSub)

}

func doesTree1HasTree2(pHead, pSub *tree.BinaryTreeNode) bool {
	// sub没有子节点了
	if pSub == nil {
		return true
	}

	if pHead == nil {
		return false
	}

	if pHead.Val != pSub.Val {
		return false
	}

	return doesTree1HasTree2(pHead.Left, pSub.Left) && doesTree1HasTree2(pHead.Right, pSub.Right)
}
