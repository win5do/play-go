package main

import (
	"fmt"
)

// 面试题7：重建二叉树
// 题目：输入某二叉树的前序遍历和中序遍历的结果，请重建出该二叉树。假设输
// 入的前序遍历和中序遍历的结果中都不含重复的数字。例如输入前序遍历序列{1,
// 2, 4, 7, 3, 5, 6, 8}和中序遍历序列{4, 7, 2, 1, 5, 3, 8, 6}，则重建出
// 图2.6所示的二叉树并输出它的头结点。

type BinaryTreeNode struct {
	Val   int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

func Construct(preOrder, inOrder []int) *BinaryTreeNode {
	if preOrder == nil || inOrder == nil || len(preOrder) == 0 {
		return nil
	}

	if len(preOrder) != len(inOrder) {
		panic("invalid input")
	}

	rootVal := preOrder[0]
	root := new(BinaryTreeNode)
	root.Val = rootVal

	rootIndex := -1
	for k, v := range inOrder {
		if v == rootVal {
			rootIndex = k
		}
	}

	if rootIndex < 0 {
		panic("invalid input")
	}

	if 1+rootIndex > len(preOrder) {
		root.Left = Construct(nil, inOrder[:rootIndex])
		root.Right = nil
	} else {
		root.Left = Construct(preOrder[1:1+rootIndex], inOrder[:rootIndex])
		root.Right = Construct(preOrder[1+rootIndex:], inOrder[rootIndex+1:])
	}

	return root
}

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
