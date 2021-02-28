package main

func hammingDistance(x int, y int) int {
	r := 0
	a := x ^ y
	for a != 0 {
		if 1&a == 1 {
			r++
		}

		a = a >> 1
	}
	return r
}
