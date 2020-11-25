package dp

import "golang.org/x/tools/go/ssa/interp/testdata/src/fmt"

func numTrees(n int) int {
	var f func(l, r int) int
	f = func(l, r int) int {
		if l == r {
			return 1
		}
		if r-l == 1 {
			return 2
		}
		count := 0
		for i := l; i <= r; i++ {
			if i == l {
				count += f(l+1, r)
			} else if i == r {
				count += f(l, r-1)
			} else {
				count += f(l, i-1)
				count += f(i+1, r)
			}
			fmt.Println(l, r, count)
		}

		return count
	}

	return f(0, n)
}
