package main

import (
	"fmt"
	"sort"
)

// 判断b是不是a的子集
// 需要考虑重复元素
func main() {
	a := []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}
	//b := []int{1, 3, 2, 1}
	b := []int{1, 3, 2, 7}

	r := checkSubset(a, b)
	r2 := checkSubset2(a, b)
	fmt.Println(r)
	fmt.Println(r2)
}

// 方法一: 可以用map储存a的元素和个数,查找b的每个元素,找到个数减一,找不到或个数为-1退出

// 方法二:
// 算法复杂度:
// m=len(a) n=len(b)
// 两次排序nlogn mlogm
// 循环次数n<=x<=m,j一直递增,最多是m
func checkSubset(a, b []int) bool {
	if len(b) > len(a) {
		return false
	}

	// 先排序
	sort.Ints(a)
	sort.Ints(b)

	i := 0
	j := 0

	for i < len(b) && j < len(a) {
		if b[i] == a[j] {
			i++
		}
		j++
	}

	if i < len(b) {
		return false
	}

	return true
}

func checkSubset2(a, b []int) bool {
	if len(b) > len(a) {
		return false
	}

	// 先排序
	sort.Ints(a)
	sort.Ints(b)

	i := 0
	j := 0

	// 上面的一次循环其实是两次循环的简写
	for i < len(b) {
		if j >= len(a) {
			return false
		}
		for j < len(a) {
			if b[i] == a[j] {
				i++
				j++
				break
			}
			j++
		}
	}

	return true
}
