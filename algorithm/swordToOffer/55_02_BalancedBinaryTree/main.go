package main

import (
	"play-go/algorithm/swordToOffer/dataStruct/tree"
)

// 面试题55（二）：平衡二叉树
// 题目：输入一棵二叉树的根结点，判断该树是不是平衡二叉树。如果某二叉树中
// 任意结点的左右子树的深度相差不超过1，那么它就是一棵平衡二叉树。

func isBalanced(pHead *tree.Tree) bool {
	depth := 0
	return isBalancedCore(pHead, &depth)
}

func isBalancedCore(pNode *tree.Tree, depth *int) bool {
	if pNode == nil {
		return true
	}

	left, right := 0, 0

	if isBalancedCore(pNode.Left, &left) &&
		isBalancedCore(pNode.Right, &right) {
		diff := left - right
		if (diff > 1) || (diff < -1) {
			return false
		}
		if diff >= 0 {
			*depth = left + 1
		} else {
			*depth = right + 1
		}

		return true
	}

	return false
}
