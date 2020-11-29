package dp

func minDistance(word1 string, word2 string) int {
	l1 := len(word1)
	l2 := len(word2)

	dp := make([][]int, l1+1)
	for i := range dp {
		dp[i] = make([]int, l2+1)
	}

	for i := 0; i <= l1; i++ {
		dp[i][0] = i
	}

	for i := 0; i <= l2; i++ {
		dp[0][i] = i
	}

	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min72(
					dp[i][j-1]+1,
					dp[i-1][j]+1,
					dp[i-1][j-1]+1)
			}
		}
	}

	return dp[l1][l2]
}

func min72(args ...int) int {
	min := args[0]
	for _, v := range args {
		if v < min {
			min = v
		}
	}

	return min
}
