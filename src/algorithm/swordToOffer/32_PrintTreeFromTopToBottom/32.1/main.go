package main

import (
	"container/list"
	"fmt"

	"playGo/src/algorithm/swordToOffer/dataStruct/tree"
)

// 面试题32（一）：不分行从上往下打印二叉树
// 题目：从上往下打印出二叉树的每个结点，同一层的结点按照从左到右的顺序打印。

func printTreeFromTopToBottom(pHead *tree.Tree) {
	if pHead == nil {
		return
	}

	quene := list.New()
	quene.PushBack(pHead)

	for quene.Len() > 0 {
		pNode := quene.Remove(quene.Front()).(*tree.Tree)

		fmt.Println(pNode.Val)
		if pNode.Left != nil {
			quene.PushBack(pNode.Left)
		}
		if pNode.Right != nil {
			quene.PushBack(pNode.Right)
		}
	}
}
