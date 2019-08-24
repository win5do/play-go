package tree

import (
	"fmt"
)

// 深度优先遍历
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

// 广度优先遍历
func BreadthFirstTravel(head *Tree) {
	if head == nil {
		return
	}

	var list []*Tree
	list = append(list, head)

	for len(list) > 0 {
		node := list[0]
		list = list[1:]
		fmt.Println(node.Val)
		if node.Left != nil {
			list = append(list, node.Left)
		}
		if node.Right != nil {
			list = append(list, node.Right)
		}
	}
}
