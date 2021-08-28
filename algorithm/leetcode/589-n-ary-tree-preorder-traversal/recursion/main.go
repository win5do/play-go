package recursion

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	var ans []int

	preorderCore(root, &ans)

	return ans
}

func preorderCore(node *Node, ans *[]int) {
	if node == nil {
		return
	}
	*ans = append(*ans, node.Val)

	for _, v := range node.Children {
		preorderCore(v, ans)
	}
}
