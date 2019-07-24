package main

import (
	"playGo/src/algorithm/swordToOffer/dataStruct/tree"
)

func getCommonParent_dfs(head, a, b *tree.BinaryTreeNode) (parent *tree.BinaryTreeNode) {
	if head == nil || a == nil || b == nil {
		return nil
	}

	var aList, bList []*tree.BinaryTreeNode

	tree.DeepFirstSearch(head, a, &aList)
	tree.DeepFirstSearch(head, b, &bList)

	return getLastCommonParent(aList, bList)
}

func getCommonParent_bfs(head, a, b *tree.BinaryTreeNode) (parent *tree.BinaryTreeNode) {
	if head == nil || a == nil || b == nil {
		return nil
	}

	var aList, bList []*tree.BinaryTreeNode

	aList, _ = tree.BreadthFirstSearch(head, a)
	bList, _ = tree.BreadthFirstSearch(head, b)

	return getLastCommonParent(aList, bList)
}

func getLastCommonParent(aList, bList []*tree.BinaryTreeNode) (common *tree.BinaryTreeNode) {
	for k, v := range aList {
		// 最后一个元素是a,b, k最多为n-2
		if len(bList)-1 > k && v == bList[k] {
			common = v
		} else {
			break
		}
	}

	return
}
