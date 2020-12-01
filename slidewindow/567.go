package slidewindow

func checkInclusion(s1 string, s2 string) bool {
	need, window := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(s1); i++ {
		need[s1[i]]++
	}
	ls1 := len(s1)
	valid := 0
	hope := len(need)
	left, right := 0, 0
	for right < len(s2) {
		vr := s2[right]
		if _, exist := need[vr]; exist {
			window[vr]++
			if window[vr] == need[vr] {
				valid++
			}
		}
		for valid == hope && (right-left+1) >= ls1 {
			if (right - left + 1) == ls1 {
				return true
			}
			vl := s2[left]
			left++
			if _, exist := need[vl]; exist {
				if window[vl] == need[vl] {
					valid--
				}
				window[vl]--
			}
		}
		right++
	}

	return false
}
