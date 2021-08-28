package main

import (
	"container/list"
	"fmt"

	"playGo/algorithm/swordToOffer/dataStruct/tree"
)

// 面试题32（二）：分行从上到下打印二叉树
// 题目：从上到下按层打印二叉树，同一层的结点按从左到右的顺序打印，每一层
// 打印到一行。

func printTree(pHead *tree.Tree) {
	if pHead == nil {
		return
	}

	quene := list.New()
	quene.PushBack(pHead)

	currentFloor := 1
	nextFloor := 0

	for quene.Len() > 0 {
		pNode := quene.Remove(quene.Front()).(*tree.Tree)

		if pNode.Left != nil {
			quene.PushBack(pNode.Left)
			nextFloor++
		}
		if pNode.Right != nil {
			quene.PushBack(pNode.Right)
			nextFloor++
		}

		fmt.Print(pNode.Val)
		currentFloor--
		if currentFloor == 0 {
			fmt.Println()
			currentFloor = nextFloor
			nextFloor = 0
		}
	}
}
