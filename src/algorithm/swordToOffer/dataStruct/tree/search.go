package tree

func DeepFirstSearch(head, node *BinaryTreeNode, list *[]*BinaryTreeNode) (found bool) {
	if head == nil || node == nil {
		return false
	}

	if head == node {
		return true
	}

	*list = append(*list, head)

	found = DeepFirstSearch(head.Left, node, list)
	if found {
		return
	}

	found = DeepFirstSearch(head.Right, node, list)
	if found {
		return
	}

	if !found {
		*list = (*list)[:len(*list)-1]
	}

	return false
}
