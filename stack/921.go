package stack

func minAddToMakeValid(S string) int {
	left, right := 0, 0
	for _, v := range S {
		if v == ')' {
			if left == 0 {
				right++
			} else {
				left--
			}
		}

		if v == '(' {
			left++
		}
	}

	return left + right
}
