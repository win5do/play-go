package main

import (
	"errors"
)

// 面试题67：把字符串转换成整数
// 题目：请你写一个函数StrToInt，实现把字符串转换成整数这个功能。当然，不
// 能使用atoi或者其他类似的库函数。

var errInvalidInput = errors.New("invalid input")

func atoi(str string) (int, error) {
	if str == "" {
		return 0, errInvalidInput
	}

	number := 0
	minus := false

	bts := []byte(str)
	if bts[0] == '-' || bts[0] == '+' {
		if bts[0] == '-' {
			minus = true
		}

		bts = bts[1:]
		if len(bts) < 1 {
			return 0, errInvalidInput
		}
	}

	for _, v := range bts {
		n := v - '0'
		if n > 9 {
			return 0, errInvalidInput
		}

		number = number*10 + int(n)
	}

	if minus {
		return -number, nil
	}

	return number, nil
}
