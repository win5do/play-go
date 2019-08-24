package main

import (
	"playGo/src/algorithm/swordToOffer/dataStruct/tree"
)

// 面试题27：二叉树的镜像
// 题目：请完成一个函数，输入一个二叉树，该函数输出它的镜像。

func mirror(pHead *tree.Tree) *tree.Tree {
	if pHead == nil {
		return nil
	}

	pHead.Left, pHead.Right = pHead.Right, pHead.Left

	pHead.Left = mirror(pHead.Left)
	pHead.Right = mirror(pHead.Right)

	return pHead
}
