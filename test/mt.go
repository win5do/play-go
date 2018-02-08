package main

import "fmt"

func main() {

}

func swap(arr []int, i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func insert(arr []int) []int {
	leng := len(arr)

	for i := 1; i < leng; i++ {
		temp := arr[i]
		j := i

		for ; j >= 1 && arr[j-1] > temp; j-- {
			arr[j] = arr[j-1]
		}

		arr[j] = temp
	}

	return arr
}
