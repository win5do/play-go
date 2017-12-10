package benchmark

import (
	"math/rand"
)

func generateData(x int) []string {
	word := "abcdefghijklmnopqrstuvwxyz"
	var data []string

	for x > 0 {
		x--
		item := ""
		for i := 0; i < 4; i++ {
			pos := rand.Intn(26)
			item += word[pos:pos+1]
		}
		data = append(data, item)
	}

	return data
}
