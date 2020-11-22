package main

import "fmt"

func stoneGameIII(stoneValue []int) string {
	subMap := make(map[int]int)
	helperMap := make(map[int]int)
	length := len(stoneValue)

	var getSum func(start int) int
	getSum = func(start int) int {
		if _, exist := subMap[start]; exist {
			return subMap[start]
		}
		sum := 0
		for i := start; i < length; i++ {
			sum += stoneValue[i]
		}
		subMap[start] = sum
		return sum
	}

	var helper func(start int) int
	helper = func(start int) int {
		if _, exist := helperMap[start]; exist {
			return helperMap[start]
		}
		if start >= length-3 {
			return getSum(start)
		}
		helperMap[start+1] = helper(start + 1)
		helperMap[start+2] = helper(start + 2)
		helperMap[start+3] = helper(start + 3)
		return getSum(start) - minIn3(helperMap[start+1], helperMap[start+2], helperMap[start+3])
	}

	rs := helper(0)*2 - getSum(0)
	if rs > 0 {
		return "Alice"
	} else if rs < 0 {
		return "Bob"
	} else {
		return "Tie"
	}
}

func minIn3(q, p, z int) int {
	if q <= p && q <= z {
		return q
	} else if p <= q && p <= z {
		return p
	} else {
		return z
	}
}

func main() {
	ls := []int{2, 3, 4, 5, 6}
	fmt.Println(stoneGameIII(ls))
}
