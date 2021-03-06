package recurse

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
	if root == nil {
		return ans
	}

	stack := []*Node{root}

	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		ans = append(ans, cur.Val)

		for i := len(cur.Children) - 1; i >= 0; i-- {
			stack = append(stack, cur.Children[i])
		}
	}

	return ans
}
