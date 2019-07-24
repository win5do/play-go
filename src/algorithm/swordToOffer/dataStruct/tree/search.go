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
func BreadthFirstSearch(head, node *BinaryTreeNode) (list []*BinaryTreeNode, found bool) {
	if head == nil || node == nil {
		return nil, false
	}

	// 二维数组记录路径
	var quene [][]*BinaryTreeNode

	quene = append(quene, []*BinaryTreeNode{head})

	for len(quene) > 0 {
		// pop
		path := quene[0]
		quene = quene[1:]

		// 对比最后一个元素
		last := path[len(path)-1]

		if last == node {
			return path, true
		}

		if last.Left != nil {
			newPath := append(path, last.Left)
			quene = append(quene, newPath)
		}

		if last.Right != nil {
			newPath := append(path, last.Right)
			quene = append(quene, newPath)
		}
	}

	return nil, false
}
