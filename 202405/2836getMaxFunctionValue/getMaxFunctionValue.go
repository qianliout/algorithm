package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Println(getMaxFunctionValue([]int{2, 0, 1}, 4))
}

func getMaxFunctionValue(receiver []int, k int64) int64 {
	type pair struct{ pa, sum int }
	n := len(receiver)
	m := bits.Len(uint(k))
	pa := make([][]pair, n)
	for i, p := range receiver {
		pa[i] = make([]pair, m)
		pa[i][0] = pair{p, p}
	}

	for i := 0; i+1 < m; i++ {
		for j := 0; j < n; j++ {
			p := pa[j][i]
			pp := pa[p.pa][i]
			// todo(还是没有级理解)
			pa[j][i+1] = pair{pp.pa, p.sum + pp.sum}
		}
	}

	ans := 0
	// 这样写也是可以的，只是不容易理解
	// for i := range receiver {
	// 	x := i
	// 	sum := i
	// 	for j := k; j > 0; j &= j - 1 {
	// 		p := pa[x][bits.TrailingZeros(uint(j))]
	// 		sum += p.sum
	// 		x = p.pa
	// 	}
	// 	ans = max(ans, sum)
	// }

	// 这样写更容易理解
	for i := range receiver {
		x := i
		sum := i
		for j := 0; j < m; j++ {
			if k>>j&1 > 0 {
				p := pa[x][j]
				sum += p.sum
				x = p.pa
			}
		}
		ans = max(ans, sum)
	}

	return int64(ans)
}
