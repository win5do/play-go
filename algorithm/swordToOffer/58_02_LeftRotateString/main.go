package main

// 面试题58（二）：左旋转字符串
// 题目：字符串的左旋转操作是把字符串前面的若干个字符转移到字符串的尾部。
// 请定义一个函数实现字符串左旋转操作的功能。比如输入字符串"abcdefg"和数
// 字2，该函数将返回左旋转2位得到的结果"cdefgab"。

func reverseString(rs []rune) {
	start := 0
	end := len(rs) - 1
	for start < end {
		rs[start], rs[end] = rs[end], rs[start]
		start++
		end--
	}
}

func leftRotateString(str string, n int) string {
	if len(str) == 0 || n >= len(str) || n <= 0 {
		return str
	}

	rs := []rune(str)
	// 先翻转前面部分，再翻转后面部分，最后翻转整个字符串
	reverseString(rs[:n])
	reverseString(rs[n:])
	reverseString(rs)

	return string(rs)
}
