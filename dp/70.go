package dp

func climbStairs(n int) int {
	if n < 3 {
		return n
	}

	pre := 2
	ppre := 1
	var res int
	for i := 3; i <= n; i++ {
		res = pre + ppre
		ppre = pre
		pre = res
	}

	return res
}
