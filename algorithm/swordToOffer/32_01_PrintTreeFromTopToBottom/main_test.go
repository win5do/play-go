package main

import (
	"testing"

	"playGo/algorithm/swordToOffer/dataStruct/tree"
)

func TestPrintTree(t *testing.T) {
	p1 := tree.NewNode(1)
	p2 := tree.NewNode(2)
	p3 := tree.NewNode(3)
	p4 := tree.NewNode(4)
	p5 := tree.NewNode(5)
	p6 := tree.NewNode(6)

	tree.ConnectTreeNodes(p1, p2, p3)
	tree.ConnectTreeNodes(p2, p4, p5)
	tree.ConnectTreeNodes(p3, nil, p6)

	printTreeFromTopToBottom(p1)
}
