package tree

import (
	"testing"

	"github.com/stretchr/testify/require"
)

/*
      a
    /   \
   b     e
  / \   / \
 c   d f   g
*/

func TestDeepFirstSearch(t *testing.T) {
	tree := ConstructTree([]string{"a", "b", "c", "d", "e", "f", "g"})
	var list []*BinaryTreeNode
	found := DeepFirstSearch(tree, tree.Left.Right, &list)
	require.True(t, found)
	for _, v := range list {
		t.Log(v.Val)
	}
}

func TestBreathedFirstSearch(t *testing.T) {
	tree := ConstructTree([]string{"a", "b", "d", "e", "c", "f", "g"})
	list, found := BreadthFirstSearch(tree, tree.Right.Left)
	require.True(t, found)
	for _, v := range list {
		t.Log(v.Val)
	}
}
