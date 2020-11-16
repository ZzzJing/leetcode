package stack

func dailyTemperatures(T []int) []int {
	length := len(T)
	stack := []int{}
	res := make([]int, length)
	for i := 0; i < length; i++ {
		for len(stack) > 0 && T[i] > T[stack[len(stack)-1]] {
			stackLen := len(stack)
			pre := stack[stackLen-1]
			stack = stack[:stackLen-1]
			res[pre] = i - pre
		}
		stack = append(stack, i)
	}

	return res
}
