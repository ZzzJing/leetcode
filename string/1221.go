package string

func balancedStringSplit(s string) int {
	l , r := 0, 0
	ans := 0
	for _, v := range s {
		if v == 'L' {
			l++
		} else {
			r++
		}

		if l == r {
			ans++
			l , r = 0, 0
		}
	}

	return ans
}
