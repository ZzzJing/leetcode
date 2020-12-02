package slidewindow

func lengthOfLongestSubstring(s string) int {
	ls := len(s)
	window := map[byte]int{}
	max := 0
	for l, r := 0, 0; r < ls; r++ {
		sr := s[r]
		window[sr]++
		for window[sr] > 1 {
			sl := s[l]
			window[sl]--
			l++
		}
		len := r - l + 1
		if len > max {
			max = len
		}
	}

	return max
}
