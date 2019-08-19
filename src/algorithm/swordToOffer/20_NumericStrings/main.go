package main

// 面试题20：表示数值的字符串
// 题目：请实现一个函数用来判断字符串是否表示数值（包括整数和小数）。例如，
// 字符串“+100”、“5e2”、“-123”、“3.1416”及“-1E-16”都表示数值，但“12e”、
// “1a3.14”、“1.2.3”、“+-5”及“12e+5.4”都不是

func isNumeric(str string) bool {
	rest := []rune(str)
	if len(rest) == 0 {
		return false
	}

	rest, result := scanInteger(rest, true)
	if !result || len(rest) == 0 {
		return result
	}

	if rest[0] == '.' {
		rest = rest[1:]
	}

	rest, result = scanInteger(rest, false)
	if !result || len(rest) == 0 {
		return result
	}

	if rest[0] == 'e' || rest[0] == 'E' {
		rest = rest[1:]
	}

	rest, result = scanInteger(rest, true)
	if !result || len(rest) == 0 {
		return result
	}

	return false
}

var intMap = make(map[rune]bool, 10)

func init() {
	for _, v := range "0123456789" {
		intMap[v] = true
	}
}

func scanInteger(rest []rune, signed bool) ([]rune, bool) {
	if len(rest) == 0 {
		return nil, false
	}

	i := 0
	result := false

	if signed && (rest[0] == '+' || rest[0] == '-') {
		i++
	}

	for i < len(rest) {
		if !intMap[rest[i]] {
			break
		} else {
			result = true
			i++
		}
	}

	return rest[i:], result
}
