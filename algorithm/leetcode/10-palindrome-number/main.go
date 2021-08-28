package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	return x == reverse(x)
}

func reverse(x int) int {
	ans := 0

	for x > 0 {
		end := x % 10
		x = x / 10
		ans = ans*10 + end
	}

	return ans
}
