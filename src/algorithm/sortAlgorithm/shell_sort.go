package sortAlgorithm

func shellSort(arr []int) []int {
	n := len(arr)
	h := 1
	for h < n/3 {
		h = 3*h + 1
	}
	for h >= 1 {
		for i := h; i < n; i++ {
			for j := i; j >= h; j -= h {
				if arr[j] < arr[j-h] {
					swap(arr, j, j-h)
				}
			}
		}

		h = h / 3
	}

	return arr
}
