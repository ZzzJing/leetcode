package dp

func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}
	dp := make([]int, n+1)
	dp[1] = 10
	for i := 2; i <= min357(n, 10); i++ {
		num := 9
		for j := 0; j < i-1; j++ {
			num = num * (9 - j)
		}
		dp[i] = num
	}

	rs := 0
	for i := 1; i <= n; i++ {
		rs += dp[i]
	}

	return rs
}

func min357(p, q int) int {
	if p < q {
		return p
	} else {
		return q
	}
}
