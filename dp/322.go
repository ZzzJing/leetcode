package dp

import "golang.org/x/tools/go/ssa/interp/testdata/src/fmt"

func coinChange(coins []int, amount int) int {
	list := make([]int, amount+1)
	list[0] = 0
	for i := 1; i < len(list); i++ {
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}
			list[i] = min(list[i], 1+list[i-coin])
		}
	}
	fmt.Println()
	if list[amount] != 0 {
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
