package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isSymmetricCore(root.Left, root.Right)
}

func isSymmetricCore(a, b *TreeNode) bool {
	if a == nil || b == nil {
		return a == b
	}

	return a.Val == b.Val && isSymmetricCore(a.Left, b.Right) && isSymmetricCore(a.Right, b.Left)
}
