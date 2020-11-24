package dp

func stoneGameIII_2(stoneValue []int) string {
	sum := 0
	n := len(stoneValue)
	dp := make([]int, n+2)

	for i := n - 1; i >= 0; i-- {
		sum += stoneValue[i]
		dp[i] = sum - minIn3_2(dp[i+1], dp[i+2], dp[i+3])
	}

	rs := dp[0]*2 - sum
	if rs > 0 {
		return "Alice"
	} else if rs < 0 {
		return "Bob"
	} else {
		return "Tie"
	}
}

func minIn3_2(q, p, z int) int {
	if q <= p && q <= z {
		return q
	} else if p <= q && p <= z {
		return p
	} else {
		return z
	}
}
