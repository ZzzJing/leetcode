package list

func validMountainArray(A []int) bool {
	length := len(A)
	left := 0
	right := length - 1

	for left+1 < length && A[left] < A[left+1] {
		left++
	}

	for right-1 > 0 && A[right-1] > A[right] {
		right--
	}

	return left > 0 && right < length-1 && left == right
}
