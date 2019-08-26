package main

import (
	"container/list"
	"fmt"

	"playGo/src/algorithm/swordToOffer/dataStruct/tree"
)

func printTreeFromTopToBottom(pHead *tree.Tree) {
	if pHead == nil {
		return
	}

	quene := list.New()
	quene.PushBack(pHead)

	for quene.Len() > 0 {
		pNode := quene.Front().Value.(*tree.Tree)
		quene.Remove(quene.Front())

			fmt.Println(pNode.Val)
		if pNode.Left != nil {
			quene.PushBack(pNode.Left)
		}
		if pNode.Right != nil {
			quene.PushBack(pNode.Right)
		}
	}
}
