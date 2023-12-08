package main

import (
	"play-go/algorithm/swordToOffer/dataStruct/tree"
)

// 面试题55（一）：二叉树的深度
// 题目：输入一棵二叉树的根结点，求该树的深度。从根结点到叶结点依次经过的
// 结点（含根、叶结点）形成树的一条路径，最长路径的长度为树的深度。

func treeDepth(pHead *tree.Tree) int {
	if pHead == nil {
		return 0
	}

	leftH := treeDepth(pHead.Left)
	rightH := treeDepth(pHead.Right)

	if leftH > rightH {
		return leftH + 1
	} else {
		return rightH + 1
	}
}
