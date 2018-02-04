package sort_algorithm

func chooseSort(arr []int) []int {
	l := len(arr)
	min := 0
	for i := 0; i < l; i++ {
		min = arr[i]
		for j := i + 1; j < l; j++ {
			if arr[j] < min {
				t := arr[j]
				arr[j] = min
				min = t
			}
		}
		arr[i] = min
	}
	return arr
}
