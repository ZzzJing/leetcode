package string

func generateParenthesis(n int) []string {
	dp := make([][]string, n+1)
	dp[0] = []string{""}
	dp[1] = []string{"()"}
	for i := 2; i<=n ; i++ {
		tmp := []string{}
		for q:=0;q<=i-1;q++ {
			tmpQ := dp[q]
			tmpP := dp[i-1-q]
			for _, vq := range tmpQ {
				for _, vp := range tmpP {
					tmp = append(tmp, "(" + vq + ")" + vp)
				}
			}
		}
		dp[i] = tmp
	}

	return dp[n]
}
