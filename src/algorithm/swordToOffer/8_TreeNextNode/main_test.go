package main

import (
	"testing"
)

func buildTree() map[string]*BinaryTreeNode {
	var a, b, c, d, e, f, g, h, i BinaryTreeNode

	a = BinaryTreeNode{
		Val:   "a",
		Left:  &b,
		Right: &c,
	}

	b = BinaryTreeNode{
		Val:    "b",
		Left:   &d,
		Right:  &e,
		Parent: &a,
	}

	c = BinaryTreeNode{
		Val:    "c",
		Left:   &f,
		Right:  &g,
		Parent: &a,
	}

	d = BinaryTreeNode{
		Val:    "d",
		Parent: &b,
	}

	e = BinaryTreeNode{
		Val:    "e",
		Left:   &h,
		Right:  &i,
		Parent: &b,
	}

	f = BinaryTreeNode{
		Val:    "f",
		Parent: &c,
	}

	g = BinaryTreeNode{
		Val:    "g",
		Parent: &c,
	}

	h = BinaryTreeNode{
		Val:    "h",
		Parent: &e,
	}

	i = BinaryTreeNode{
		Val:    "i",
		Parent: &e,
	}

	return map[string]*BinaryTreeNode{
		"a": &a,
		"b": &b,
		"c": &c,
		"d": &d,
		"e": &e,
		"f": &f,
		"g": &g,
		"h": &h,
		"i": &i,
	}
}

func TestGetNext(t *testing.T) {
	arr := buildTree()
	t.Log(GetNext(arr["d"]))
	t.Log(GetNext(arr["b"]))
	t.Log(GetNext(arr["h"]))
	t.Log(GetNext(arr["e"]))
	t.Log(GetNext(arr["i"]))
	t.Log(GetNext(arr["a"]))
	t.Log(GetNext(arr["f"]))
	t.Log(GetNext(arr["c"]))
	t.Log(GetNext(arr["g"]))
}
