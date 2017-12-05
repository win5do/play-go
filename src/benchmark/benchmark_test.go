package benchmark

import (
	"sort"
	"testing"
)

var data []string

func sortDuplicate(arr []string) []string {
	var result []string
	for i, leng := 0, len(arr); i < leng; i++ {
		if i == 0 {
			result = append(result, arr[i])
			continue
		}

		if (arr[i] != arr[i-1]) {
			result = append(result, arr[i])
		}
	}

	return result
}

func mapDuplicate(arr []string) []string {
	result := make(map[string]bool)
	for i, leng := 0, len(arr); i < leng; i++ {
		if i == 0 {
			result[arr[i]] = true
			continue
		}

		if !result[arr[i]] {
			result[arr[i]] = true
		}
	}

	ret := []string{}

	for k := range result {
		ret = append(ret, k)
	}

	return ret
}


func BenchmarkTest1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sort.Strings(data)
		sortDuplicate(data)
	}
}

func BenchmarkTest2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mapDuplicate(data)
	}
}

func init() {
	data = generateData(1000000)
}