package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var ppath []*TreeNode
	var qpath []*TreeNode
	dfs(root, p, &ppath)
	dfs(root, q, &qpath)

	return getCommon(ppath, qpath)
}

func getCommon(a, b []*TreeNode) *TreeNode {
	ans := 0

	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] == b[i] {
			ans = i
		} else {
			break
		}
	}

	return a[ans]
}

func dfs(node, in *TreeNode, path *[]*TreeNode) bool {
	if node == nil {
		return false
	}

	*path = append(*path, node)

	if node == in {
		return true
	}

	r := dfs(node.Left, in, path) || dfs(node.Right, in, path)

	if !r {
		*path = (*path)[:len(*path)-1]
	}

	return r
}
