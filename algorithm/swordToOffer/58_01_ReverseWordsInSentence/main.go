package main

// 面试题58（一）：翻转单词顺序
// 题目：输入一个英文句子，翻转句子中单词的顺序，但单词内字符的顺序不变。
// 为简单起见，标点符号和普通字母一样处理。例如输入字符串"I am a student."，
// 则输出"student. a am I"。

// 需要将字符串先转rune再操作，否则不支持中文(不过中文也不用空格分词)
func reverseString(rs []rune) {
	start := 0
	end := len(rs) - 1
	for start < end {
		rs[start], rs[end] = rs[end], rs[start]
		start++
		end--
	}
}

func reverseSentence(str string) string {
	if len(str) == 0 {
		return ""
	}

	rs := []rune(str)
	reverseString(rs)

	start := 0
	end := 0

	for end < len(rs) {
		if end == len(rs)-1 {
			reverseString(rs[start : end+1])
		} else if rs[end] == ' ' {
			reverseString(rs[start:end])
			start = end + 1
		}

		end++
	}

	return string(rs)
}
