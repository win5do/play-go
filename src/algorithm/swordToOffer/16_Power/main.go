package main

// --- 循环实现 ---
func power_loop(base float64, exponent int) float64 {
	if exponent == 0 {
		return 1
	}

	var result float64 = base

	var absExponent int
	if exponent < 0 {
		absExponent = -exponent
	} else {
		absExponent = exponent
	}

	for absExponent > 1 {
		// 用位与判断奇数更高效，只有第一次会是奇数
		if absExponent&1 == 1 {
			result *= base
		}

		absExponent = absExponent >> 1
		result *= result
	}

	// 负数
	if exponent < 0 {
		result = 1 / result
	}

	return result
}

// --- 递归实现 ---
func power_recurse(base float64, exponent int) float64 {
	var absExponent int
	if exponent < 0 {
		absExponent = -exponent
	} else {
		absExponent = exponent
	}

	result := powerUnsigned(base, absExponent)

	// 负数
	if exponent < 0 {
		result = 1 / result
	}

	return result
}

func powerUnsigned(base float64, exponent int) float64 {
	if exponent == 0 {
		return 1
	}
	if exponent == 1 {
		return base
	}

	result := powerUnsigned(base, exponent>>1)
	result *= result

	// 奇数，用位与判断更高效
	if exponent&1 == 1 {
		result *= base
	}

	return result
}
