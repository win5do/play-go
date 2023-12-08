package array

import "play-go/algorithm/sortAlgorithm"

func Partition(arr []int, start, end int) int {
	return sortAlgorithm.Partition_Lomuto(arr, start, end)
}
