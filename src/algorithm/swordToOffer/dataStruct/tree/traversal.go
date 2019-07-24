package tree

import (
	"fmt"
)

// 深度优先遍历
func PreOrderTraversal(head *BinaryTreeNode) {
	if head == nil {
		return
	}

	fmt.Println(head.Val)
	PreOrderTraversal(head.Left)
	PreOrderTraversal(head.Right)
}

func InOrderTraversal(head *BinaryTreeNode) {
	if head == nil {
		return
	}

	InOrderTraversal(head.Left)
	fmt.Println(head.Val)
	InOrderTraversal(head.Right)
}

func PostOrderTraversal(head *BinaryTreeNode) {
	if head == nil {
		return
	}

	PostOrderTraversal(head.Left)
	PostOrderTraversal(head.Right)
	fmt.Println(head.Val)
}

// 广度优先遍历
func BreadthFirstTravel(head *BinaryTreeNode) {
	if head == nil {
		return
	}

	var list []*BinaryTreeNode
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
