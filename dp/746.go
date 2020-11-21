package dp

func minCostClimbingStairs(cost []int) int {
	length := len(cost)
	if length == 0 || length == 1 {
		return 0
	}
	res := 0
	pre := 0
	ppre := 0
	for i := 2; i <= length; i++ {
		res = min(pre+cost[i-1], ppre+cost[i-2])
		ppre = pre
		pre = res
	}

	return res
}

func min(p, q int) int {
	if p < q {
		return p
	} else {
		return q
	}
}
