package main

import (
	"fmt"
)

// 面试题15：二进制中1的个数
// 题目：请实现一个函数，输入一个整数，输出该数二进制表示中1的个数。例如
// 把9表示成二进制是1001，有2位是1。因此如果输入9，该函数输出2。

// 二进制储存的是补码，负数一直右移最终会变成-1，-1补码为-11111111，-1右移还是-1，所以负数会死循环
func NumberOf1_unsign(n int) (count int) {
	if n < 0 {
		panic("nonsupport negative integer")
	}

	for n != 0 {
		if n&1 != 0 {
			count++
		}
		n = n >> 1
		fmt.Printf("%064b\n", n)
	}
	return count
}

func NumberOf1_leftShift(n int) (count int) {
	flag := 1
	// 每次左移一位，整数最多64位，最后会溢出变成0
	for flag != 0 {
		if n&flag != 0 {
			count++
		}
		flag = flag << 1
		fmt.Printf("%064b\n", flag)
	}
	return count
}

func NumberOf1_and(n int) (count int) {
	for n != 0 {
		count++
		// 把一个整数减1再和原整数做按位与运算，会把该整数最右边的1变为0
		n = (n - 1) & n
	}
	return count
}
