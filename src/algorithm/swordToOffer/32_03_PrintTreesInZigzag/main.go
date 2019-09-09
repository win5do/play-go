package main

import (
	"container/list"
	"fmt"

	"playGo/src/algorithm/swordToOffer/dataStruct/tree"
)

// 面试题32（三）：之字形打印二叉树
// 题目：请实现一个函数按照之字形顺序打印二叉树，即第一行按照从左到右的顺
// 序打印，第二层按照从右到左的顺序打印，第三行再按照从左到右的顺序打印，
// 其他行以此类推。

func printTree(pHead *tree.Tree) {
	if pHead == nil {
		return
	}

	var quene *list.List
	// 奇偶栈
	quene1 := list.New()
	quene2 := list.New()

	quene = quene1
	quene.PushBack(pHead)

	even := false // 当前是奇数层还是偶数层

	for quene.Len() > 0 {
		// pop
		pNode := quene.Remove(quene.Back()).(*tree.Tree)

		if even {
			// 偶数层子节点进奇数栈
			if pNode.Right != nil {
				quene1.PushBack(pNode.Right)
			}
			if pNode.Left != nil {
				quene1.PushBack(pNode.Left)
			}
		} else {
			// 奇数层子节点进偶数栈
			if pNode.Left != nil {
				quene2.PushBack(pNode.Left)
			}
			if pNode.Right != nil {
				quene2.PushBack(pNode.Right)
			}
		}

		fmt.Print(pNode.Val)
		if quene.Len() == 0 {
			fmt.Println()
			even = !even
			if even {
				quene = quene2
			} else {
				quene = quene1
			}
		}
	}
}
