package main

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	var ans [][]int
	var path []int
	pathCore(root, &path, &ans)
	return toStr(ans)
}

func pathCore(node *TreeNode, path *[]int, ans *[][]int) {
	if node == nil {
		return
	}

	*path = append(*path, node.Val)

	defer func() {
		*path = (*path)[:len(*path)-1]
	}()

	if node.Left == nil && node.Right == nil {
		r := make([]int, len(*path))
		copy(r, *path)
		*ans = append(*ans, r)
		return
	}

	pathCore(node.Left, path, ans)
	pathCore(node.Right, path, ans)
}

func toStr(in [][]int) []string {
	r := make([]string, len(in))
	for i := range in {
		var buf strings.Builder
		for k, v := range in[i] {
			if k != 0 {
				buf.WriteString("->")
			}

			buf.WriteString(strconv.Itoa(v))
		}

		r[i] = buf.String()
	}

	return r
}
