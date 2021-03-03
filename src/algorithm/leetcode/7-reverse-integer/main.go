package main

func reverse(x int) int {
	var negative bool
	if x < 0 {
		negative = true
		x = abs(x)
	}

	ans := 0

	for x > 0 {
		end := x % 10
		x = x / 10
		ans = ans*10 + end
	}

	if negative {
		ans = -ans
	}

	return check(ans)
}

func abs(in int) int {
	if in < 0 {
		return -in
	}
	return in
}

func check(in int) int {
	if in > 1<<31-1 || in < -1<<31 {
		return 0
	}

	return in
}
