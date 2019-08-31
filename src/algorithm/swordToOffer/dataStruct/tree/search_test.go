package tree

import (
	"testing"

	"github.com/stretchr/testify/require"
)

/*
      1
    /   \
   2     5
  / \   / \
 3   4 6   7
*/

func TestDeepFirstSearch(t *testing.T) {
	tree := ConstructTree([]int{1, 2, 3, 4, 5, 6, 7})
	var list []*Tree
	found := DeepFirstSearch(tree, tree.Left.Right, &list)
	require.True(t, found)
	for _, v := range list {
		t.Log(v.Val)
	}
}

func TestBreathedFirstSearch(t *testing.T) {
	tree := ConstructTree([]int{1, 2, 3, 4, 5, 6, 7})
	list, found := BreadthFirstSearch(tree, tree.Right.Left)
	require.True(t, found)
	for _, v := range list {
		t.Log(v.Val)
	}
}
