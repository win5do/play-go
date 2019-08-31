package tree

import (
	"fmt"
	"testing"
)

func TestConstructTree(t *testing.T) {
	cs := []struct {
		input []int
	}{
		{
			[]int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			[]int{1, 2},
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			[]int{1},
		},
		{
			[]int{-1},
		},
		{
			[]int{1, 2, 3, -1, 4, -1, 5},
		},
	}

	for _, c := range cs {
		tree := ConstructTree(c.input)
		fmt.Println("preOrder:")
		PreOrderTraversal(tree)
		fmt.Println("inOrder")
		InOrderTraversal(tree)
		fmt.Println("-----")
	}
}

func TestPostOrderTraversal(t *testing.T) {
	tree := ConstructTree([]int{1, 2, 3, 4, 5, 6, 7})
	PostOrderTraversal(tree)
}

func TestBreadthFirstTravel(t *testing.T) {
	tree := ConstructTree([]int{1, 2, 3, 4, 5, 6, 7})
	BreadthFirstTravel(tree)
}
