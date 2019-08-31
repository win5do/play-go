package main

import (
	"fmt"
	"testing"

	"playGo/src/algorithm/swordToOffer/dataStruct/tree"
)

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
	for _, v := range r {
		for _, w := range v {
			fmt.Println(w.Val)
		}
		fmt.Println("---")
	}
}

func TestFindPath_nil(t *testing.T) {
	p1 := tree.NewTreeNode(1)
	r := findPath(p1, 1)
	for _, v := range r {
		for _, w := range v {
			fmt.Println(w.Val)
		}
		fmt.Println("---")
	}
}

func TestFindPath_one(t *testing.T) {
	r := findPath(nil, 0)
	for _, v := range r {
		for _, w := range v {
			fmt.Println(w.Val)
		}
		fmt.Println("---")
	}
}
