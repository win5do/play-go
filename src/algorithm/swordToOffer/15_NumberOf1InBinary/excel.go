package main

var alphabetMap = make(map[rune]int, 26)

func init() {
	for i, v := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		alphabetMap[v] = i + 1
	}
}

func excelNumber(str string) int {
	rs := []rune(str)
	result := 0
	base := 1
	bit := len(alphabetMap) // *26

	for i := len(str) - 1; i >= 0; i-- {
		n, ok := alphabetMap[rs[i]]
		if !ok {
			panic("invalid input")
		}

		result += n * base
		base *= bit // 26^n
	}

	return result
}
