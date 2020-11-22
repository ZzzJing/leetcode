package dp

func countVowelStrings(n int) int {
	if n == 0 {
		return 0
	}
	last := []int{1, 1, 1, 1, 1}
	for i := 2; i <= n; i++ {
		tmp := make([]int, 5)
		copy(tmp, last)
		for j := 0; j < 5; j++ {
			for k := j + 1; k < 5; k++ {
				last[k] += tmp[j]
			}
		}
	}
	res := 0
	for _, v := range last {
		res += v
	}
	return res
}
