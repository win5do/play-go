package tree

import (
	"fmt"
)

type BinaryTreeNode struct {
	Val   string
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

// 前序遍历顺序完全二叉树
func ConstructTree(items []string) (tree *BinaryTreeNode) {
	n := len(items)
	if n == 0 {
		return nil
	}
	head := items[0]
	if head == "#" {
		return nil
	}

	tree = new(BinaryTreeNode)
	tree.Val = head

	i := n/2 + 1 // 左右均分，不相等时左比右多1
	if i > n {
		return
	}

	left := items[1:i]
	right := items[i:]

	if len(left) > 0 {
		tree.Left = ConstructTree(left)
	}
	if len(right) > 0 {
		tree.Right = ConstructTree(right)
	}

	return tree
}

// func DeepFirstTraversal(pHead, pNode *BinaryTreeNode) (list []*BinaryTreeNode, found bool) {
// }

// 深度优先遍历
func PreOrderTraversal(root *BinaryTreeNode) {
	if root == nil {
		return
	}

	fmt.Println(root.Val)
	PreOrderTraversal(root.Left)
	PreOrderTraversal(root.Right)
}

func InOrderTraversal(root *BinaryTreeNode) {
	if root == nil {
		return
	}

	InOrderTraversal(root.Left)
	fmt.Println(root.Val)
	InOrderTraversal(root.Right)
}

func PostOrderTraversal(root *BinaryTreeNode) {
	if root == nil {
		return
	}

	PostOrderTraversal(root.Left)
	PostOrderTraversal(root.Right)
	fmt.Println(root.Val)
}

// 广度优先遍历
func BreadthFirstTravel(root *BinaryTreeNode) {
	if root == nil {
		return
	}

	var list []*BinaryTreeNode
	list = append(list, root)

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
