package main

func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}

	return n&(n-1) == 0
}

func isPowerOfTwoLogN(n int) bool {
	count := 0

	for n != 0 && count < 2 {
		if n&1 == 1 {
			count++
		}

		n >>= 1
	}

	return count == 1
}
