package main

import "fmt"

func main() {

}

func swap(arr []int, i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func bubble(arr []int) []int {
	leng := len(arr)
	for i := leng - 1; i > 0; i-- {
		for j := 0; j < i-1; j++ {
			if arr[j] > arr[j+1] {
				swap(arr, j, j+1)
			}
		}
	}
	return arr
}
