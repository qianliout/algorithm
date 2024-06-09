package main

import (
	"math"
)

func main() {

}

func maxScore(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	n, m := len(grid), len(grid[0])

	pre := make([][]int, n+1) // i,j 之前的最小值
	for i := range pre {
		pre[i] = make([]int, m+1)
	}

	for i := 0; i < len(pre); i++ {
		pre[i][0] = math.MaxInt
	}
	for i := 1; i < len(pre[0]); i++ {
		pre[0][i] = math.MaxInt
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			pre[i+1][j+1] = min(pre[i][j+1], pre[i+1][j], grid[i][j])
		}
	}
	ans := math.MinInt

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			ans = max(ans, grid[i][j]-min(pre[i][j+1], pre[i+1][j]))
		}
	}

	return ans
}
