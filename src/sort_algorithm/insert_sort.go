package sort_algorithm

func insert(arr []int) []int {
	l := len(arr)
	for i := 1; i < l; i++ {
		for j := 0; j < i; j++ {
			if arr[i] < arr[j] {
				t := arr[i]
				arr[i] = arr[j]
				arr[j] = t
			}
		}
	}
	return arr
}
