package list

func minEatingSpeed(piles []int, H int) int {
	right := getMax(piles) + 1
	left := 1
	for left < right {
		mid := (left + right) / 2
		if canEat(piles, mid, H) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func canEat(list []int, speed, H int) bool {
	time := 0
	for i:=0;i<len(list);i++{
		add := 0
		if list[i] % speed > 0 {
			add = 1
		}

		time += list[i]/speed + add
	}
	return time <= H
}

func getMax(list []int) int {
	max := list[0]
	for i:=0;i<len(list);i++ {
		if list[i] > max {
			max = list[i]
		}
	}

	return max
}
