package main

import (
	"playGo/src/algorithm/swordToOffer/dataStruct/tree"
)

func getCommonParent_list(head, a, b *tree.BinaryTreeNode) (parent *tree.BinaryTreeNode) {
	if head == nil || a == nil || b == nil {
		return nil
	}

	var aList, bList []*tree.BinaryTreeNode

	tree.DeepFirstSearch(head, a, &aList)
	tree.DeepFirstSearch(head, b, &bList)

	var common *tree.BinaryTreeNode
	for k, v := range aList {
		if len(bList) > k && v == bList[k] {
			common = v
		} else {
			break
		}
	}

	return common
}
