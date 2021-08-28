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

	lstack := []*TreeNode{root.Left}
	rstack := []*TreeNode{root.Right}

	for {
		lleng := len(lstack)
		rleng := len(rstack)
		if lleng == 0 || rleng == 0 {
			return lleng == rleng
		}

		ltop := lstack[lleng-1]
		lstack = lstack[:lleng-1]
		rtop := rstack[rleng-1]
		rstack = rstack[:rleng-1]

		if ltop == nil || rtop == nil {
			if ltop != rtop {
				return false
			}
			continue
		}

		if ltop.Val != rtop.Val {
			return false
		}

		lstack = append(lstack, ltop.Right, ltop.Left)
		rstack = append(rstack, rtop.Left, rtop.Right)
	}
}
