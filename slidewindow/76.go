package slidewindow

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
	ansL, ansR := -1, -1
	left, right := 0, 0
	for right < len(s) {
		if _, exist := window[s[right]]; exist {
			window[s[right]]++
		} else {
			window[s[right]] = 1
		}
		right++
		for check() {
			ansL = left
			ansR = right
			window[s[left]]--
			left++
		}

	}
}
