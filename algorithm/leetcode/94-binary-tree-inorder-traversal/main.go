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

func inorderTraversal(root *TreeNode) []int {
	list := []int{}

	cur := root
	var stack []*TreeNode
	for cur != nil {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}

		for {
			l := len(stack)
			if l <= 0 {
				break
			}

			top := stack[l-1]
			stack = stack[:l-1]
			list = append(list, top.Val)

			if top.Right != nil {
				cur = top.Right
				break
			}
		}
	}

	return list
}
