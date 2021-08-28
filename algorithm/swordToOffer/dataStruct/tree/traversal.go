package tree

import (
	"container/list"
	"fmt"
)

// 前序遍历 根-左-右
// 深度优先遍历
func PreOrderTraversal(head *Tree) {
	if head == nil {
		return
	}

	fmt.Println(head.Val)
	PreOrderTraversal(head.Left)
	PreOrderTraversal(head.Right)
}

// 中序遍历 左-根-右
func InOrderTraversal(head *Tree) {
	if head == nil {
		return
	}

	InOrderTraversal(head.Left)
	fmt.Println(head.Val)
	InOrderTraversal(head.Right)
}

// 后序遍历 左-右-根
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

	stack := list.New() // 栈
	stack.PushBack(head)

	for stack.Len() > 0 {
		pNode := stack.Remove(stack.Back()).(Tree)
		fmt.Println(pNode.Val)

		if pNode.Right != nil {
			stack.PushBack(pNode.Right)
		}
		if pNode.Left != nil {
			stack.PushBack(pNode.Left)
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
