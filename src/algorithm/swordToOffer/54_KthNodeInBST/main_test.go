package main

import (
	"testing"

	"playGo/src/algorithm/swordToOffer/dataStruct/tree"

	"github.com/stretchr/testify/assert"
)

func TestFunc(t *testing.T) {
	pHead := tree.ConstructTree([]int{5, 3, 2, 4, 7, 6, 8})
	tree.InOrderTraversal(pHead)
	r := kthNodeInBST(pHead, 3)
	assert.Equal(t, 4, r.Val)
}