package main

import (
	"playGo/src/algorithm/swordToOffer/dataStruct/stack"
)

// 面试题31：栈的压入、弹出序列
// 题目：输入两个整数序列，第一个序列表示栈的压入顺序，请判断第二个序列是
// 否为该栈的弹出顺序。假设压入栈的所有数字均不相等。例如序列1、2、3、4、
// 5是某栈的压栈序列，序列4、5、3、2、1是该压栈序列对应的一个弹出序列，但
// 4、3、5、1、2就不可能是该压栈序列的弹出序列。

func isPopOrder(pushOrder, popOrder []int) bool {
	if len(pushOrder) == 0 || len(popOrder) == 0 || len(pushOrder) != len(popOrder) {
		return false
	}

	length := len(pushOrder)
	stk := new(stack.Stack)

	i, j := 0, 0

	for i < length {
		for j < length {
			if stk.Size() > 0 && stk.Top() == popOrder[i] {
				break
			}
			stk.Push(pushOrder[j])
			j++
		}
		if stk.Size() == 0 || stk.Pop() != popOrder[i] {
			return false
		}
		i++
	}

	// if j != length-1 && i != j {
	// 	return false
	// }

	return true
}
