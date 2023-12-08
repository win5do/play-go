package main

import (
	"play-go/algorithm/swordToOffer/dataStruct/tree"
)

// 面试题28：对称的二叉树
// 题目：请实现一个函数，用来判断一棵二叉树是不是对称的。如果一棵二叉树和
// 它的镜像一样，那么它是对称的。

// --- method 1 ---
func isSymmetricalTree(pHead *tree.Tree) bool {
	// 自己和自己对比
	return symmetricalCore(pHead, pHead)
}

// tree1的前序遍历（中左右）应该和tree2的对称前序遍历（中右左）相等
func symmetricalCore(p1, p2 *tree.Tree) bool {
	if p1 == nil && p2 == nil {
		return true
	}

	if p1 == nil || p2 == nil {
		return false
	}

	if p1.Val != p2.Val {
		return false
	}

	return symmetricalCore(p1.Left, p2.Right) && symmetricalCore(p1.Right, p2.Left)
}

// --- method 2 ---
// 先求树的镜像，再遍历对比两个树的节点。
// 不如上面的方法。
