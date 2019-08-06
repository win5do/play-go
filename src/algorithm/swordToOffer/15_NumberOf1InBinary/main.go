package main

import (
	"fmt"
)

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
