package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type nodeLevel struct {
	node  *TreeNode
	level int
}

func levelOrder(root *TreeNode) [][]int {
	var ans [][]int
	if root == nil {
		return ans
	}

	q := []*nodeLevel{
		{
			node:  root,
			level: 0,
		},
	}

	for len(q) > 0 {
		item := q[0]
		q = q[1:]

		if len(ans) == item.level {
			ans = append(ans, []int{})
		}
		ans[item.level] = append(ans[item.level], item.node.Val)

		nextLevel := item.level + 1
		if item.node.Left != nil {
			q = append(q, &nodeLevel{
				node:  item.node.Left,
				level: nextLevel,
			})
		}

		if item.node.Right != nil {
			q = append(q, &nodeLevel{
				node:  item.node.Right,
				level: nextLevel,
			})
		}
	}

	return ans
}
