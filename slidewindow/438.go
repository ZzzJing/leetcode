package slidewindow

import "fmt"

func findAnagrams(s string, p string) []int {
	need, window := map[byte]int{}, map[byte]int{}
	ls := len(s)
	lp := len(p)
	for i := 0; i < lp; i++ {
		need[p[i]]++
	}
	hope := len(need)
	result := []int{}
	vaild := 0

	for l, r := 0, 0; r < ls; r++ {
		sr := s[r]
		if _, exist := need[sr]; exist {
			window[sr]++
			if need[sr] == window[sr] {
				vaild++
			}
		}

		for vaild == hope && (r-l+1) >= lp {
			fmt.Println(l, r, r-l+1, lp)
			if (r - l + 1) == lp {
				result = append(result, l)
			}
			sl := s[l]
			l++
			if _, exist := need[sl]; exist {
				if window[sl] == need[sl] {
					vaild--
				}
				window[sl]--
			}
		}
	}

	return result
}
