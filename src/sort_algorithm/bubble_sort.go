package sort_algorithm

func bubble(arr []interface{}) []interface{} {
	l := len(arr)
	for i := 0; i < l-1; i++ {
		if compare(arr[i], arr[i+1]) {
			t := arr[i]
			arr[i] = arr[i+1]
			arr[i+1] = t
		}
	}
	return arr
}

func compare(a interface{}, b interface{}) bool {
	ia, ioka := a.(int)
	ib, iokb := b.(int)
	sa, soka := a.(string)
	sb, sokb := b.(string)

	if ioka && iokb {
		return ia > ib
	} else if soka && sokb {
		return sa > sb
	} else if ioka && sokb {
		return string(ia) > sb
	} else if soka && iokb {
		return sa > string(ib)
	} else {
		return false
	}
}
