package main

import (
	"playGo/algorithm/swordToOffer/dataStruct/tree"
)

// 面试题68：树中两个结点的最低公共祖先
// 题目：输入两个树结点，求它们的最低公共祖先。

// 思路：用DFS或BFS找到根节点到该节点的路径，
// 再找到两个路径中最后一个相同节点

// --- dfs ---
func getCommonParent_dfs(head, a, b *tree.Tree) (parent *tree.Tree) {
	if head == nil || a == nil || b == nil {
		return nil
	}

	var aList, bList []*tree.Tree

	tree.DeepFirstSearch(head, a, &aList)
	tree.DeepFirstSearch(head, b, &bList)

	return getLastCommonParent(aList, bList)
}

// --- bfs ---
func getCommonParent_bfs(head, a, b *tree.Tree) (parent *tree.Tree) {
	if head == nil || a == nil || b == nil {
		return nil
	}

	var aList, bList []*tree.Tree

	aList, _ = tree.BreadthFirstSearch(head, a)
	bList, _ = tree.BreadthFirstSearch(head, b)

	return getLastCommonParent(aList, bList)
}

// --- common ---
func getLastCommonParent(aList, bList []*tree.Tree) (common *tree.Tree) {
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
