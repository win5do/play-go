package main

// 面试题19：正则表达式匹配
// 题目：请实现一个函数用来匹配包含'.'和'*'的正则表达式。模式中的字符'.'
// 表示任意一个字符，而'*'表示它前面的字符可以出现任意次（含0次）。在本题
// 中，匹配是指字符串的所有字符匹配整个模式。例如，字符串"aaa"与模式"a.a"
// 和"ab*ac*a"匹配，但与"aa.a"及"ab*a"均不匹配

func match(str, pattern string) bool {
	if len(str) == 0 && len(pattern) == 0 {
		return true
	}

	if len(str) != 0 && len(pattern) == 0 {
		return false
	}

	if len(pattern) > 1 && pattern[1] == '*' {
		// pattern第二个字符为*

		if len(str) > 0 && (str[0] == pattern[0] || pattern[0] == '.') {
			// 第一个字符匹配有两种情况，可继续匹配.*，也可以跳过.*
			return match(str[1:], pattern) || match(str[1:], pattern[2:])
		} else {
			// 不匹配，模式跳过.*
			// len(str) == 0但pattern还剩若干.*的情况也成立
			return match(str, pattern[2:])
		}
	}

	if len(str) > 0 && (str[0] == pattern[0] || pattern[0] == '.') {
		return match(str[1:], pattern[1:])
	}

	return false
}
