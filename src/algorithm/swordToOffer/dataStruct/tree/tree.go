package tree

type Tree struct {
	Val   int
	Left  *Tree
	Right *Tree
}

// 此常数表示nil
const NilNode = -99999

// 使用完全二叉树的中序遍历构造树
func ConstructTree(items []int) (tree *Tree) {
	n := len(items)
	if n == 0 {
		return nil
	}
	head := items[0]
	if head == NilNode {
		return nil
	}

	tree = new(Tree)
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

func NewTreeNode(val int) *Tree {
	return &Tree{
		Val: val,
	}
}

func ConnectTreeNodes(pParent, pLeft, pRight *Tree) {
	pParent.Left = pLeft
	pParent.Right = pRight
}
