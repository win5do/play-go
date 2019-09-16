package main

import (
	"log"
)

// 面试题5：替换空格
// 题目：请实现一个函数，把字符串中的每个空格替换成"%20"。例如输入“We are happy.”，
// 则输出“We%20are%20happy.”。

func replaceBlank(input string) string {
	originRune := []rune(input)

	originLength := 0
	countBlank := 0

	for _, v := range originRune {
		originLength++
		if v == ' ' {
			countBlank++
		}
	}

	originIndex := originLength - 1
	newIndex := originIndex + 2*countBlank
	newRune := append(originRune, make([]rune, 2*countBlank)...)
	log.Print("originRune length: ", len(originRune))
	log.Print("newRune length: ", len(newRune))

	for originIndex >= 0 && newIndex > originIndex {
		if newRune[originIndex] == ' ' {
			newRune[newIndex] = '0'
			newRune[newIndex-1] = '2'
			newRune[newIndex-2] = '%'

			newIndex = newIndex - 2
		} else {
			newRune[newIndex] = newRune[originIndex]
		}
		originIndex--
		newIndex--
	}

	return string(newRune)
}
