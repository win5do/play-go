package main

func reverseWords(s string) string {
	b := []byte(s)

	start := 0
	end := 0

	for i := range b {
		if b[i] == ' ' {
			end = i
		} else if i == len(b)-1 {
			end = i + 1
		}

		if end > start {
			reverseString(b[start:end])
			start = end + 1
		}
	}

	return string(b)
}

func reverseString(s []byte) {
	leng := len(s)
	for l, r := 0, leng-1; l < r; l, r = l+1, r-1 {
		swap(s, l, r)
	}
}

func swap(arr []byte, a, b int) {
	arr[a], arr[b] = arr[b], arr[a]
}
