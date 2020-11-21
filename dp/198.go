package dp

func rob(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return nums[0]
	}
	if length == 2 {
		return max(nums[0], nums[1])
	}
	pre := max(nums[0], nums[1])
	ppre := nums[0]
	res := 0
	for i := 2; i < length; i++ {
		res = max(pre, ppre+nums[i])
		ppre = pre
		pre = res
	}

	return res
}

func max(p, q int) int {
	if p > q {
		return p
	}

	return q
}
