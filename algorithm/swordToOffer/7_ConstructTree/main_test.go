package main

import (
	"fmt"
	"testing"
)

func TestConstruct(t *testing.T) {
	cs := []struct {
		preOrder, inOrder []int
	}{
		{
			[]int{1, 2, 4, 7, 3, 5, 6, 8},
			[]int{4, 7, 2, 1, 5, 3, 8, 6},
		},
		{
			nil,
			nil,
		},
		{
			[]int{},
			[]int{},
		},
		{
			[]int{1},
			[]int{1},
		},
		{
			// all left
			[]int{1, 2, 3, 4, 5, 6, 7, 8},
			[]int{8, 7, 6, 5, 4, 3, 2, 1},
		},
		{
			// all right
			[]int{1, 2, 3, 4, 5, 6, 7, 8},
			[]int{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			// not match
			[]int{1, 2, 4, 7, 3, 5, 6, 8},
			[]int{1, 7, 2, 4, 5, 3, 8, 6},
		},
	}

	for k, v := range cs {
		root := Construct(v.preOrder, v.inOrder)
		fmt.Println("case", k)
		PreOrderTraversal(root)
		fmt.Println("-----")
		InOrderTraversal(root)
		fmt.Println("-----")
		PostOrderTraversal(root)
	}
}
