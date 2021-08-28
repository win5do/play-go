package main

import "sort"

func removeDuplication_map(arr []string) []string {
	set := make(map[string]struct{}, len(arr))

	j := 0
	for _, v := range arr {
		_, ok := set[v]
		if ok {
			continue
		}
		set[v] = struct{}{}
		arr[j] = v
		j++
	}

	return arr[:j]
}

func removeDuplication_sort(arr []string) []string {
	sort.Strings(arr)

	length := len(arr)
	if length == 0 {
		return arr
	}

	j := 0
	for i := 1; i < length; i++ {
		if arr[i] != arr[j] {
			j++
			if j < i {
				swap(arr, i, j)
			}
		}
	}

	return arr[:j+1]
}

func swap(arr []string, a, b int) {
	arr[a], arr[b] = arr[b], arr[a]
}
