package sort_algorithm

func quickSort(arr []int) []int {
	return quickRecurse(arr, 0, len(arr)-1)
}

func quickRecurse(arr []int, start, end int) [] int {
	if start > end-1 {
		return arr
	}

	m := arr[start]
	lp := start
	rp := end

	for lp < rp {
		// 从右往左找到第一个比m小的下标
		for arr[rp] >= m && lp < rp {
			rp--
		}

		// 从左往右找到第一个比m大的下标
		for arr[lp] <= m && lp < rp {
			lp++
		}

		arr[rp], arr[lp] = arr[lp], arr[rp]
	}

	// lp == rp
	arr[lp], arr[start] = m, arr[lp]

	quickRecurse(arr, start, lp-1)
	quickRecurse(arr, lp+1, end)

	return arr
}

func quickSort2(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	lp, rp := 0, len(arr)-1

	for lp < rp {
		// arr[0]为基准数 arr[lp]一直等于arr[0]
		if arr[lp+1] > arr[lp] {
			// 移到后面
			arr[lp+1], arr[rp] = arr[rp], arr[lp+1]
			rp--
		} else {
			// 向前移
			arr[lp+1], arr[lp] = arr[lp], arr[lp+1]
			lp++
		}
	}

	quickSort2(arr[:lp])
	quickSort2(arr[lp+1:])

	return arr
}
