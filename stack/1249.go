package stack

import "strings"

func minRemoveToMakeValid(s string) string {
	length := len(s)
	rs := ""
	left := 0
	tmp := []rune(s)
	for i := 0; i < length; i++ {
		if tmp[i] == '(' {
			left++
		} else if tmp[i] == ')' {
			if left == 0 {
				tmp[i] = '0'
			} else {
				left--
			}
		}
	}

	rs = strings.ReplaceAll(string(tmp), "0", "")

	if left > 0 {
		for j := len(rs) - 1; j >= 0; j-- {
			if rs[j:j+1] == "(" && left > 0 {
				left--
				rs = rs[:j] + rs[j+1:]
			}

		}
	}
	return rs
}
