package main

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

	var pNext *BinaryTreeNode

	if pNode.Right != nil {
		pNext = pNode.Right
		for pNext.Left != nil {
			pNext = pNext.Left
		}
		return pNext
	} else if pNode.Parent != nil {
		pCurrent := pNode
		pNext = pNode.Parent
		for pNext != nil && pNext.Right == pCurrent {
			pCurrent = pNext
			pNext = pCurrent.Parent
		}

		return pNext
	}

	return nil
}
