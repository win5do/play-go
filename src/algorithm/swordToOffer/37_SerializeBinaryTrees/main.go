package main

import (
	"strconv"
	"strings"

	"playGo/src/algorithm/swordToOffer/dataStruct/tree"
)

// 面试题37：序列化二叉树
// 题目：请实现两个函数，分别用来序列化和反序列化二叉树。

func serialize(pHead *tree.Tree) string {
	if pHead == nil {
		return ""
	}
	var str string
	serializeCore(pHead, &str)
	return strings.TrimSuffix(str, ",")
}

func serializeCore(pNode *tree.Tree, str *string) {
	if pNode == nil {
		*str += "$,"
		return
	} else {
		val := strconv.Itoa(pNode.Val)
		*str += val + ","
	}

	serializeCore(pNode.Left, str)
	serializeCore(pNode.Right, str)
}

func deserialize(str string) *tree.Tree {
	if str == "" {
		return nil
	}

	strs := strings.Split(str, ",")

	return deserializeCore(&strs)
}

func deserializeCore(strs *[]string) *tree.Tree {
	if len(*strs) == 0 {
		return nil
	}

	s0 := (*strs)[0]
	*strs = (*strs)[1:]
	if s0 == "$" {
		return nil
	}

	val, err := strconv.Atoi(s0)
	if err != nil {
		panic("invalid serialize")
	}
	pNode := tree.NewNode(val)

	pNode.Left = deserializeCore(strs)
	pNode.Right = deserializeCore(strs)

	return pNode
}
