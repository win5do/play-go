package main

import (
	"fmt"
	"testing"

	"playGo/src/algorithm/swordToOffer/dataStruct/tree"
)

func printResult(r [][]*tree.Tree) {
	for _, v := range r {
		for _, w := range v {
			fmt.Println(w.Val)
		}
		fmt.Println("---")
	}
}

func TestFindPath(t *testing.T) {
	p1 := tree.NewTreeNode(1)
	p2 := tree.NewTreeNode(2)
	p3 := tree.NewTreeNode(3)
	p4 := tree.NewTreeNode(4)
	p5 := tree.NewTreeNode(5)
	p6 := tree.NewTreeNode(6)
	p_2 := tree.NewTreeNode(-2)

	tree.ConnectTreeNodes(p1, p2, p3)
	tree.ConnectTreeNodes(p2, p4, p5)
	tree.ConnectTreeNodes(p3, p4, p6)
	tree.ConnectTreeNodes(p6, p_2, nil) // 负值

	r := findPath(p1, 8)
	printResult(r)

	// nil
	r = findPath(nil, 0)
	printResult(r)

	// one node
	r = findPath(tree.NewTreeNode(1), 1)
	printResult(r)
}
