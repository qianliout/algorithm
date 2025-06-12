package main

import (
	"fmt"
)

func main() {
	fmt.Println(numSquares(12))

}

func numSquares(n int) int {
	var dfs func(i int) int
	mem := make([]int, n+10)

	dfs = func(i int) int {
		if i <= 1 {
			return i
		}
		ans := i
		if mem[i] > 0 {
			return mem[i]
		}
		for k := 1; k*k <= i; k++ {
			ans = min(ans, dfs(i-k*k)+1)
		}
		mem[i] = ans
		return ans
	}
	ans := dfs(n)
	return ans
}
