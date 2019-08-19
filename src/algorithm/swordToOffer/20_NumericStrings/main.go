package main

var intMap = make(map[rune]bool, 10)

func init() {
	for _, v := range "0123456789" {
		intMap[v] = true
	}
}

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
