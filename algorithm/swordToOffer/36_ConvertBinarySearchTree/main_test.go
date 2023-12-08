package main

import (
	"fmt"
	"testing"

	"play-go/algorithm/swordToOffer/dataStruct/tree"
)

func traversal(pHead *tree.Tree) {
	pNode := pHead
	for pNode != nil {
		fmt.Printf("val: %v, prev: %v, next: %v\n", pNode.Val, pNode.Left, pNode.Right)
		pNode = pNode.Right
	}
}

func TestConvert(t *testing.T) {
	cs := []struct {
		input []int
	}{
		{input: []int{}},
		{input: []int{10, 6, 4, 8, 14, 12, 16}},
		{input: []int{10}},
		{input: []int{10, 6, 4}},
		{input: []int{10, 6}},
		{input: []int{10, tree.NilNode, 14}},
		{input: []int{10, 6, 4, tree.NilNode, tree.NilNode}},
		{input: []int{10, tree.NilNode, tree.NilNode, tree.NilNode, 14, tree.NilNode, 16}},
	}

	for i, v := range cs {
		fmt.Println("case:", i)
		r := tree.ConstructTree(v.input)
		r = covert(r)
		traversal(r)
		fmt.Println("---")
	}
}
