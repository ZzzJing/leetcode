package dp

func maxSubArray(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return nums[0]
	}
	max := 0
	tmp := 0
	maxOne := nums[0]
	for i := 0; i < length; i++ {
		tmp = tmp + nums[i]

		if tmp > max {
			max = tmp
		}

		if tmp < 0 {
			tmp = 0
		}

		if nums[i] > maxOne {
			maxOne = nums[i]
		}
	}

	if max != 0 {
		return max
	} else {
		return maxOne
	}
}
