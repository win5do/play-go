package main

func reverseString(s []byte) {
	leng := len(s)
	for l, r := 0, leng-1; l < r; l, r = l+1, r-1 {
		swap(s, l, r)
	}
}

func swap(arr []byte, a, b int) {
	arr[a], arr[b] = arr[b], arr[a]
}
