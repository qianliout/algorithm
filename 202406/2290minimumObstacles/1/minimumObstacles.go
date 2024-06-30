package main

import (
	"math"
)

func main() {

}

// 这样写有问题
func minimumObstacles(grid [][]int) int {
	inf := math.MaxInt / 100
	m, n := len(grid), len(grid[0])
	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = inf
		}
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for i := 1; i < n; i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}

	// 这样写只能是 从上到下从左到右
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			ch := grid[i][j]
			d := inf
			for _, dir := range dirs {
				x, y := i+dir[0], j+dir[1]
				if in(m, n, x, y) {
					d = min(d, dp[x][y])
				}
			}
			if ch == 1 {
				d++
			}
			dp[i][j] = d
		}
	}

	return dp[m-1][n-1]
}

func in(m, n, c, r int) bool {
	if c < 0 || r < 0 {
		return false
	}
	if c >= m || r >= n {
		return false
	}
	return true
}
