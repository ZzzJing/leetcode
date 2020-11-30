package slidewindow

import "golang.org/x/tools/go/ssa/interp/testdata/src/fmt"

func minWindow(s string, t string) string {
	need, window := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	check := func() bool {
		for k, v := range need {
			if window[k] < v {
				return false
			}
		}

		return true
	}
	length := -1
	ansL := -1
	for left, right := 0, 0; right < len(s); right++ {
		if _, exist := window[s[right]]; exist {
			window[s[right]]++
		} else {
			window[s[right]] = 1
		}
		for check() && left <= right {
			tmpL := right - left + 1
			if length == -1 {
				length = tmpL
				ansL = left
			} else {
				if tmpL < length {
					ansL = left
					length = tmpL
				}
			}
			window[s[left]]--
			left++
		}
	}

	if ansL == -1 {
		return ""
	}
	fmt.Println(ansL, length)
	return s[ansL : ansL+length]
}
