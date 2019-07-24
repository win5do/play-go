package tree

import (
	"fmt"
	"testing"
)

func TestConstructTree(t *testing.T) {
	cs := []struct {
		input []string
	}{
		{
			[]string{"a", "b", "c", "d", "e", "f", "g"},
		},
		{
			[]string{"a", "b"},
		},
		{
			[]string{"a", "b", "c", "d", "e", "f", "g", "h"},
		},
		{
			[]string{"a"},
		},
		{
			[]string{"#"},
		},
		{
			[]string{"a", "b", "c", "#", "d", "#", "e"},
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

func TestBreadthFirstTravel(t *testing.T) {
	tree := ConstructTree([]string{"a", "b", "c", "d", "e", "f"})
	BreadthFirstTravel(tree)
}
