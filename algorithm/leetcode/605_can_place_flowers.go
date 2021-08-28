package leetcode

func canPlaceFlowers(flowerbed []int, n int) bool {
	count := 0

	prev := -1
	leng := len(flowerbed)

	for i := 0; i < leng; i++ {
		if flowerbed[i] == 1 {
			if prev < 0 {
				// flowerbed[0] == 0 时特殊处理
				count += i / 2
			} else {
				count += (i - prev - 2) / 2
			}
			prev = i
		} else {
			if i == leng-1 {
				// flowerbed[leng-1] == 0 时特殊处理
				if prev < 0 {
					count += (leng + 1) / 2
				} else {
					count += (leng - prev - 1) / 2
				}
			}
		}
	}

	return n <= count
}
