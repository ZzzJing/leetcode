package dp

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	big := 999999
	list := make([]int, amount+1)
	list[0] = 0
	for i := 1; i < len(list); i++ {
		list[i] = big
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}
			list[i] = min(list[i], 1+list[i-coin])
		}
	}

	if list[amount] != big {
		return list[amount]
	} else {
		return -1
	}
}

func min(p, q int) int {
	if p < q {
		return p
	} else {
		return q
	}
}
