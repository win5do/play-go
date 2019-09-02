package main

import (
	"container/list"

	"playGo/src/algorithm/swordToOffer/dataStruct/tree"
)

// 面试题37：序列化二叉树
// 题目：请实现两个函数，分别用来序列化和反序列化二叉树。

func serialize(pHead *tree.Tree) []int {
	items := make([]int, 0)

	quene := list.New()
	quene.PushBack(pHead)

	for quene.Len() > 0 {
		pNode := quene.Remove(quene.Back()).(tree.Tree)
		items = append(items, pNode.Val)

		if pNode.Right != nil {
			quene.PushBack(pNode.Right)
			if pNode.Left == nil {
				quene.PushBack(tree.NewNode(tree.NilNode))
			}
		}
		if pNode.Left != nil {
			quene.PushBack(pNode.Left)
			if pNode.Right == nil {
				quene.PushBack(tree.NewNode(tree.NilNode))
			}
		}
	}

	return items
}

func deserialize(items []int) *tree.Tree {
	return tree.ConstructTree(items)
}
