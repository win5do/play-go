package main

// 面试题33：二叉搜索树的后序遍历序列
// 题目：输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历的结果。
// 如果是则返回true，否则返回false。假设输入的数组的任意两个数字都互不相同。

func verifySequenceOfBST(seq []int) bool {
	length := len(seq)

	if length == 0 {
		return false
	}

	if length == 1 {
		return true
	}

	root := seq[length-1]
	i := 0
	for i < length-1 {
		if seq[i] > root {
			break
		}
		i++
	}

	for j := i; j < length-1; j++ {
		if seq[j] < root {
			return false
		}
	}

	left, right := true, true

	if i > 0 {
		left = verifySequenceOfBST(seq[:i])
	}

	if i < length-1 {
		right = verifySequenceOfBST(seq[i : length-1])
	}

	return left && right
}
