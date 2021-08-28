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

func isBalanced(root *TreeNode) bool {
	return height(root) >= 0
}

func height(node *TreeNode) int {
	if node == nil {
		return 0
	}

	lh := height(node.Left)
	rh := height(node.Right)
	if lh == -1 || rh == -1 || abs(lh-rh) > 1 {
		return -1
	}

	return max(lh, rh) + 1
}

func abs(in int) int {
	if in < 0 {
		return -in
	}
	return in
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
