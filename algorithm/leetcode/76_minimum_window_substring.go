package leetcode

func minWindow(s string, t string) string {
	expect := make(map[uint8]int)

	for i := 0; i < len(t); i++ {
		expect[t[i]] += 1
	}

	result := ""
	l, r := 0, 0
	set := make(map[uint8]int)

	for ; r < len(s) && l <= r; r++ {
		if _, ok := expect[s[r]]; !ok {
			if _, ok := expect[s[l]]; !ok {
				l++
			}
			continue
		}

		set[s[r]] += 1
		if !check(expect, set) {
			continue
		}

		result = updateResult(s[l:r+1], result)

		// 移动 l
		for l < r {
			set[s[l]] -= 1
			l++

			if _, ok := expect[s[l]]; !ok {
				continue
			}

			if !check(expect, set) {
				break
			}

			result = updateResult(s[l:r+1], result)
		}
	}

	return result
}

func updateResult(tmp, result string) string {
	if result == "" || len(tmp) < len(result) {
		result = tmp
	}
	return result
}

func check(expect, set map[uint8]int) bool {
	for k, v := range expect {
		if set[k] < v {
			return false
		}
	}

	return true
}
