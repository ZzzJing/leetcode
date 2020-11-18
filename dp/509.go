package dp

func fib(N int) int {
	if N == 1 || N == 2 {
		return 1
	}

	pre := 1
	ppre := 1
	now := 0
	for i := 3; i <= N; i++ {
		now = pre + ppre
		ppre = pre
		pre = now
	}

	return now
}
