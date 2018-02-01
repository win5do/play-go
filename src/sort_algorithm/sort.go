package sort_algorithm

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
