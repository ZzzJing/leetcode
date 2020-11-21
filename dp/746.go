package dp

func minCostClimbingStairs(cost []int) int {
	length := len(cost)
	if length == 0 || length == 1 {
		return 0
	}
	res := make([]int, length+1)
	res[0] = 0
	res[1] = 0
	for i := 2; i <= length; i++ {
		res[i] = min(res[i-1]+cost[i-1], res[i-2]+cost[i-2])
	}

	return res[length]
}

func min(p, q int) int {
	if p < q {
		return p
	} else {
		return q
	}
}
