package string

func maxDepth(s string) int {
	dep:=0
	left := 0
	for i:=0;i<len(s);i++ {
 		if s[i] == ')' {
			left--
			continue
		}
		if s[i] == '(' {
			left++
			if left > dep {
				dep = left
			}
		}
	}

	return dep
}
