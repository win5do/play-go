package main

// 面试题8：二叉树的下一个结点
// 题目：给定一棵二叉树和其中的一个结点，如何找出中序遍历顺序的下一个结点？
// 树中的结点除了有两个分别指向左右子结点的指针以外，还有一个指向父结点的指针。

type BinaryTreeNode struct {
	Val    string
	Left   *BinaryTreeNode
	Right  *BinaryTreeNode
	Parent *BinaryTreeNode
}

func GetNext(pNode *BinaryTreeNode) *BinaryTreeNode {
	if pNode == nil {
		return nil
	}

	if pNode.Right != nil {
		pNext := pNode.Right
		for pNext.Left != nil {
			pNext = pNext.Left
		}
		return pNext
	} else if pNode.Parent != nil {
		pCurrent := pNode
		pNext := pNode.Parent
		for pNext != nil && pNext.Right == pCurrent {
			pCurrent = pNext
			pNext = pCurrent.Parent
		}

		return pNext
	}

	return nil
}
