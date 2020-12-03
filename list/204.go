package list

func countPrimes(n int) int {
	list:= make([]int, n)
	ans := 0
	for i:=2;i<n;i++ {
		if list[i] == 0 {
			ans++
			for j:=i*i;j<n;j+=i {
				list[j] = 1
			}
		}
	}

	return ans
}