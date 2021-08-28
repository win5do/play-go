package main

import (
	"math/rand"
	"strings"
	"testing"
)

func generateData(x int) []string {
	word := "abcdefghijklmnopqrstuvwxyz"
	var data []string

	prev := ""
	n := 0

	for n < x {
		if n&1 == 0 {
			// even number
			var item strings.Builder
			for i := 0; i < 2; i++ {
				pos := rand.Intn(26)
				item.WriteString(word[pos : pos+1])
			}
			data = append(data, item.String())
			prev = item.String()
		} else {
			data = append(data, prev)
		}

		n++
	}

	return data
}

func copyData(in []string) []string {
	r := make([]string, len(in))
	copy(r, in)
	return r
}

func checkDuplication(arr []string) bool {
	set := make(map[string]struct{}, len(arr))
	for _, v := range arr {
		_, ok := set[v]
		if ok {
			return false
		}
		set[v] = struct{}{}
	}

	return true
}

func TestRm(t *testing.T) {
	tt := []int{0, 1, 2, 1000000}

	for _, v := range tt {
		t.Run("map", func(t *testing.T) {
			checkDuplication(removeDuplication_map(generateData(v)))
		})

		t.Run("sort", func(t *testing.T) {
			checkDuplication(removeDuplication_sort(generateData(v)))
		})
	}
}

func BenchmarkRm(b *testing.B) {
	testdata := generateData(1000000)

	b.Run("map", func(b *testing.B) {
		in := copyData(testdata)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			removeDuplication_map(in)
		}
	})

	b.Run("sort", func(b *testing.B) {
		in := copyData(testdata)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			removeDuplication_sort(in)
		}
	})
}
