package main

import (
	"fmt"
)

func main() {
	fmt.Println(countSubmatrices([][]int{{7, 6, 3}, {6, 6, 1}}, 18))
}

func countSubmatrices(grid [][]int, k int) int {
	n, m := len(grid), len(grid[0])
	pre := make([][]int, n+1)
	for i := range pre {
		pre[i] = make([]int, m+1)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			pre[i+1][j+1] = pre[i+1][j] + pre[i][j+1] - pre[i][j] + grid[i][j]
		}
	}
	ans := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if pre[i][j] <= k {
				ans++
			}
		}
	}
	return ans
}
