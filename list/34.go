package list

func searchRange(nums []int, target int) []int {
	rs := []int{-1, -1}
	findNum := -1
	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if nums[mid] == target {
			findNum = mid
			break
		} else if nums[mid] < target {
			low = mid + 1
			continue
		} else {
			high = mid - 1
		}
	}

	if findNum != -1 {
		rs[0] = findNum
		rs[1] = findNum
		for i := findNum; i >= 0; i-- {
			if nums[i] == target {
				rs[0] = i
			} else {
				break
			}
		}

		for i := findNum; i < len(nums); i++ {
			if nums[i] == target {
				rs[1] = i
			} else {
				break
			}
		}
	}
	return rs
}
