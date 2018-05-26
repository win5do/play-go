package sort_algorithm

func mergeSort(arr []int) []int {
	leng := len(arr)
	if leng < 2 {
		return arr
	}

	m := leng / 2
	left := arr[:m]
	right := arr[m:]

	left = mergeSort(left)
	right = mergeSort(right)

	return merge(left, right)
}

func merge(left, right []int) []int {
	i, j := 0, 0
	temp := []int{}

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			temp = append(temp, left[i])
			i++
		} else {
			temp = append(temp, right[j])
			j++
		}
	}

	temp = append(temp, left[i:]...)
	temp = append(temp, right[j:]...)

	return temp
}

// 迭代
func mergeSort2(arr []int) []int {
	leng := len(arr)
	k := 1
	temp := []int{};
	for k < leng {
		temp = []int{}
		for i := 0; i < leng-1; i += 2 * k {
			if i+k >= leng {
				temp = append(temp, arr[i:]...)
			} else if i+2*k >= leng {
				temp = append(temp, merge(arr[i:i+k], arr[i+k:])...)
			} else {
				temp = append(temp, merge(arr[i:i+k], arr[i+k:i+2*k])...)
			}
		}
		arr = temp
		k *= 2
	}
	return arr
}
