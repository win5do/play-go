package sort_algorithm

func insertSort(arr []int) []int {
	l := len(arr)
	for i := 1; i < l; i++ {
		t := arr[i]
		j := i - 1
		for j >= 0 {
			if t < arr[j] {
				arr[j+1] = arr[j]
				j--
			} else {
				break
			}
		}
		arr[j+1] = t
	}
	return arr
}

func insertSort2(arr []int) []int {
	l := len(arr)
	for i := 1; i < l; i++ {
		for j := i; j >= 1; j-- {
			if arr[j] < arr[j-1] {
				swap(arr, j, j-1)
			} else {
				break
			}
		}
	}
	return arr
}
