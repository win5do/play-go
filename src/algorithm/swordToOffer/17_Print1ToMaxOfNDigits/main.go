package main

import (
	"fmt"
	"strconv"
)

// n很大时会溢出
func print1ToN_onlySmall(n int) {
	x := 1
	for i := 0; i < n; i++ {
		n = x * 10
	}

	for i := 0; i < x; i++ {
		fmt.Println(i)
	}
}

// 在string上模拟数字加法
func print1ToN_string(n int) {
	if n <= 0 {
		return
	}

	numberStr := ""
	for i := 0; i < n; i++ {
		numberStr += "0"
	}

	for {
		var ok bool
		numberStr, ok = increment(numberStr)
		if ok {
			return
		}
		printNumberString(numberStr)
	}
}

func increment(str string) (newStr string, overflow bool) {
	nTakeOver := 0

	rs := []rune(str)

	overflow = true // 默认是false

	for i := len(rs) - 1; i >= 0; i-- {
		num := int(rs[i] - '0')
		inc := num + nTakeOver

		if i == len(rs)-1 {
			inc++
		}

		if inc >= 10 {
			if i == 0 {
				return "", true
			}

			nTakeOver = 1
			inc -= 10
		} else {
			overflow = false
		}

		rs[i] = rune(strconv.Itoa(inc)[0])

		if !overflow {
			break
		}
	}
	newStr = string(rs)
	return newStr, false
}

func printNumberString(str string) {
	var start int

	for i, v := range str {
		if v != '0' {
			start = i
			break
		}

		if i == len(str)-1 {
			return
		}
	}

	fmt.Println(str[start:])
}
