package tree

type BinaryTreeNode struct {
	Val   string
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

// 前序遍历顺序完全二叉树
func ConstructTree(items []string) (tree *BinaryTreeNode) {
	n := len(items)
	if n == 0 {
		return nil
	}
	head := items[0]
	if head == "#" {
		return nil
	}

	tree = new(BinaryTreeNode)
	tree.Val = head

	i := n/2 + 1 // 左右均分，不相等时左比右多1
	if i > n {
		return
	}

	left := items[1:i]
	right := items[i:]

	if len(left) > 0 {
		tree.Left = ConstructTree(left)
	}
	if len(right) > 0 {
		tree.Right = ConstructTree(right)
	}

	return tree
}
