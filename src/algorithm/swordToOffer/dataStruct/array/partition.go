package array

import "playGo/src/algorithm/sortAlgorithm"

func Partition(arr []int, start, end int) int {
	return sortAlgorithm.Partition_Lomuto(arr, start, end)
}
