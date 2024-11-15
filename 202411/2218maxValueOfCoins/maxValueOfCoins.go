package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxValueOfCoins([][]int{{1, 100, 3}, {7, 8, 9}}, 2))
}

func maxValueOfCoins(piles [][]int, k int) int {
	n := len(piles)
	sum := make([][]int, n)
	for i := range sum {
		sum[i] = gen(piles[i])
	}
	mem := make([][]int, n)
	for i := range mem {
		mem[i] = make([]int, 2000+5)
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}
	var dfs func(i int, j int) int
	dfs = func(i int, j int) int {
		if i < 0 {
			return 0
		}
		if j <= 0 {
			return 0
		}
		if mem[i][j] != -1 {
			return mem[i][j]
		}
		ans := 0
		for m := 0; m <= min(len(piles[i]), j); m++ {
			ans = max(ans, dfs(i-1, j-m)+sum[i][m])
		}
		mem[i][j] = ans
		return ans
	}
	ans := dfs(n-1, k)
	return ans
}

func gen(p []int) []int {
	ans := make([]int, len(p)+1)
	for i, ch := range p {
		ans[i+1] = ans[i] + ch
	}
	return ans
}
