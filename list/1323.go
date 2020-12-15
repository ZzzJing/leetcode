package list

func maximum69Number(num int) int {
	nums := []int{}
	for num > 0 {
		nums = append(nums, num%10)
		num = num / 10
	}

	isReplace := false
	res := 0
	for i := len(nums) - 1; i >= 0; i-- {
		if !isReplace && nums[i] == 6 {
			nums[i] = 9
			isReplace = true
		}
		res = res*10 + nums[i]
	}

	return res
}
