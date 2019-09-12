package main

import (
	"errors"
	"regexp"
)

// 面试题67：把字符串转换成整数
// 题目：请你写一个函数StrToInt，实现把字符串转换成整数这个功能。当然，不
// 能使用atoi或者其他类似的库函数。

var errInvalidInput = errors.New("invalid input")

func atoi(str string) (int, error) {
	if str == "" {
		return 0, errInvalidInput
	}

	ok, err := regexp.MatchString(`[+-]?[0-9]+$`, str)
	if !ok || err != nil {
		return 0, errInvalidInput
	}

	number := 0
	minus := false

	for i, v := range []rune(str) {
		if i == 0 {
			if v == '-' {
				minus = true
				continue
			}

			if v == '+' {
				continue
			}
		}
		number = number*10 + int(v-'0')
	}

	if minus {
		return -number, nil
	}

	return number, nil
}
