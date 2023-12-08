package main

import (
	"testing"

	"play-go/algorithm/swordToOffer/dataStruct/list"

	"github.com/stretchr/testify/assert"
)

func TestFunc(t *testing.T) {
	p1 := list.CreateListNode(1)
	p2 := list.CreateListNode(2)
	p3 := list.CreateListNode(3)
	p4 := list.CreateListNode(4)
	p5 := list.CreateListNode(5)
	p6 := list.CreateListNode(6)
	p7 := list.CreateListNode(7)

	list.ConnectListNodes(p1, p2)
	list.ConnectListNodes(p2, p3)
	list.ConnectListNodes(p3, p6)
	list.ConnectListNodes(p6, p7)
	list.ConnectListNodes(p4, p5)
	list.ConnectListNodes(p5, p6)

	r := findFirstCommonNode(p1, p4)
	assert.Equal(t, p6, r)
}
