package main

import (
	"testing"

	"playGo/src/algorithm/swordToOffer/dataStruct/tree"

	"github.com/stretchr/testify/require"
)

func TestGetCommonParent_list(t *testing.T) {
	head := tree.ConstructTree([]string{"a", "b", "c", "d", "e", "f", "g"})

	cs := []struct {
		a      *tree.Tree
		b      *tree.Tree
		expect *tree.Tree
	}{
		{
			a:      head.Left,
			b:      head.Right,
			expect: head,
		},
		{
			a:      head.Left,
			b:      new(tree.Tree),
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
		// 相同节点
		{
			a:      head,
			b:      head,
			expect: nil,
		},
	}

	for _, c := range cs {
		r := getCommonParent_dfs(head, c.a, c.b)
		t.Log(r)
		require.Equal(t, c.expect, r)

		r = getCommonParent_bfs(head, c.a, c.b)
		t.Log(r)
		require.Equal(t, c.expect, r)
	}
}
