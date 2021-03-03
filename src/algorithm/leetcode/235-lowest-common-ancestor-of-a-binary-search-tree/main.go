package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	return getLastCommon(search(root, p), search(root, q))
}

func search(root, in *TreeNode) []*TreeNode {
	var r []*TreeNode

	cur := root

	for cur != nil {
		r = append(r, cur)

		if in == cur {
			return r
		}

		if in.Val > cur.Val {
			cur = cur.Right
		} else {
			cur = cur.Left
		}
	}

	return r
}

func getLastCommon(a, b []*TreeNode) *TreeNode {
	var r *TreeNode

	for i := range a {
		if len(b) < i+1 || a[i] != b[i] {
			break
		}

		r = a[i]
	}

	return r
}
