package main

import (
	"testing"

	"playGo/src/algorithm/swordToOffer/dataStruct/tree"
)

func TestPrintTree(t *testing.T) {
	p1 := tree.NewTreeNode(1)
	p2 := tree.NewTreeNode(2)
	p3 := tree.NewTreeNode(3)
	p4 := tree.NewTreeNode(4)
	p5 := tree.NewTreeNode(5)
	p6 := tree.NewTreeNode(6)

	tree.ConnectTreeNodes(p1, p2, p3)
	tree.ConnectTreeNodes(p2, p4, p5)
	tree.ConnectTreeNodes(p3, nil, p6)

	printTree(p1)
}
