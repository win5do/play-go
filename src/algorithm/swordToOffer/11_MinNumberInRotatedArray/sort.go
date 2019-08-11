package main

// 对公司所有员工年龄排序，有几万名员工。
// 利用年龄在0~99范围的特性。
func sortAges(ages []int) {
	timesOfAge := make([]int, 100)

	for _, v := range ages {
		if v < 0 || v > 99 {
			panic("age out of range")
		}
		timesOfAge[v]++
	}

	index := 0
	for k, v := range timesOfAge {
		for v > 0 {
			ages[index] = k
			v--
			index++
		}
	}
}
