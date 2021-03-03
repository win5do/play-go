package main

func longestCommonPrefix(strs []string) string {
	ans := ""
OUT:
	for i := 0; ; i++ {
		tmp := ""

		if len(strs) == 0 {
			break OUT
		}

		for _, v := range strs {
			if len(v) < i+1 {
				break OUT
			}

			s := string(v[i])

			if tmp == "" {
				tmp = s
			} else if s != tmp {
				break OUT
			}
		}

		ans += tmp
	}

	return ans
}
