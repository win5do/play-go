package main

import (
	"playGo/src/algorithm/swordToOffer/dataStruct/tree"
)

// 面试题34：二叉树中和为某一值的路径
// 题目：输入一棵二叉树和一个整数，打印出二叉树中结点值的和为输入整数的所
// 有路径。从树的根结点开始往下一直到叶结点所经过的结点形成一条路径。

func findPath(pHead *tree.Tree, number int) [][]*tree.Tree {
	if pHead == nil {
		return nil
	}

	result := make([][]*tree.Tree, 0)

	findPathCore(pHead, number, &result, nil)

	return result
}

func findPathCore(pNode *tree.Tree, number int, result *[][]*tree.Tree, path []*tree.Tree) {
	if pNode == nil {
		return
	}

	if path == nil {
		path = make([]*tree.Tree, 0)
	}

	path = append(path, pNode)

	total := 0
	for _, v := range path {
		total += v.Val
	}

	if total == number {
		*result = append(*result, path)
	}

	findPathCore(pNode.Left, number, result, path)
	findPathCore(pNode.Right, number, result, path)
}
