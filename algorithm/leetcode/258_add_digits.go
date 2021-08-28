package leetcode

func addDigits(num int) int {
	for num > 9 {
		num = num%10 + num/10
	}
	return num
}

func addDigits_2(num int) int {
	return (num-1)%9 + 1
}
