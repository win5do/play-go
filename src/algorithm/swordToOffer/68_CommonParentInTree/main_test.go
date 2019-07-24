package main

import (
	"testing"

	"playGo/src/algorithm/swordToOffer/dataStruct/tree"

	"github.com/stretchr/testify/require"
)

func TestGetCommonParent_list(t *testing.T) {
	head := tree.ConstructTree([]string{"a", "b", "c", "d", "e", "f", "g"})

	cs := []struct {
		a      *tree.BinaryTreeNode
		b      *tree.BinaryTreeNode
		expect *tree.BinaryTreeNode
	}{
		{
			a:      head.Left,
			b:      head.Right,
			expect: head,
		},
		{
			a:      head.Left,
			b:      new(tree.BinaryTreeNode),
			expect: nil,
		},
		{
			a:      head.Left.Left,
			b:      head.Right.Right,
			expect: head,
		},
		{
			a:      head.Left.Left,
			b:      head.Left.Right,
			expect: head.Left,
		},
	}

	for _, c := range cs {
		r := getCommonParent_list(head, c.a, c.b)
		t.Log(r)
		require.Equal(t, c.expect, r)
	}
}
