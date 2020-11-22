package dp

func waysToStep(n int) int {
	if n < 3 {
		return n
	}

	pre := 2
	ppre := 1
	pppre := 1
	res := 0
	for i := 3; i <= n; i++ {
		res = (pre + ppre + pppre) % 1000000007

		pppre = ppre
		ppre = pre
		pre = res
	}

	return res
}
