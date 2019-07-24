package tree

// DFS
func DeepFirstSearch(head, node *BinaryTreeNode, list *[]*BinaryTreeNode) (found bool) {
	if head == nil || node == nil {
		return false
	}

	*list = append(*list, head)

	// list包含node
	if head == node {
		return true
	}

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

// BFS
func BreadthFirstSearch(head, node *BinaryTreeNode, list *[]*BinaryTreeNode) (found bool) {
	if head == nil || node == nil {
		return false
	}

	var quene []*BinaryTreeNode

	quene = append(quene, head)

	level := 0

	for len(quene) > 0 {
		pop := quene[0]
		quene = quene[1:]

		// list为搜索路径
		*list = append(*list, pop)
		if pop == node {
			return true
		}

		if pop.Left != nil {
			quene = append(quene, pop.Left)
		}

		if pop.Right != nil {
			quene = append(quene, pop.Right)
		}

		level++
	}

	return false
}
