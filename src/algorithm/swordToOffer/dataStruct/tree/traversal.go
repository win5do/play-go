package tree

import (
	"container/list"
	"fmt"
)

// 前序遍历 深度优先遍历
func PreOrderTraversal(head *Tree) {
	if head == nil {
		return
	}

	fmt.Println(head.Val)
	PreOrderTraversal(head.Left)
	PreOrderTraversal(head.Right)
}

func InOrderTraversal(head *Tree) {
	if head == nil {
		return
	}

	InOrderTraversal(head.Left)
	fmt.Println(head.Val)
	InOrderTraversal(head.Right)
}

func PostOrderTraversal(head *Tree) {
	if head == nil {
		return
	}

	PostOrderTraversal(head.Left)
	PostOrderTraversal(head.Right)
	fmt.Println(head.Val)
}

// 深度优先遍历_循环实现
func PreOrderTraversal_loop(head *Tree) {
	if head == nil {
		return
	}

	quene := list.New() // 栈
	quene.PushBack(head)

	for quene.Len() > 0 {
		pNode := quene.Remove(quene.Back()).(Tree)
		fmt.Println(pNode.Val)

		if pNode.Right != nil {
			quene.PushBack(pNode.Right)
		}
		if pNode.Left != nil {
			quene.PushBack(pNode.Left)
		}
	}
}

// 广度优先遍历
func BreadthFirstTravel(head *Tree) {
	if head == nil {
		return
	}

	var list []*Tree // 队列
	list = append(list, head)

	for len(list) > 0 {
		pNode := list[0]
		list = list[1:]
		fmt.Println(pNode.Val)
		if pNode.Left != nil {
			list = append(list, pNode.Left)
		}
		if pNode.Right != nil {
			list = append(list, pNode.Right)
		}
	}
}
