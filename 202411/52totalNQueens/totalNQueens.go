package main

import (
	"fmt"
)

func main() {
	fmt.Println(totalNQueens(4))
	fmt.Println(totalNQueens(9))
}

func totalNQueens(n int) int {
	grid := make([][]int, n)
	for i := range grid {
		grid[i] = make([]int, n)
	}
	ans := 0
	var dfs func(x int)
	dfs = func(x int) {
		// 找完了
		if x == n {
			ans++
			return
		}
		for i := 0; i < n; i++ {
			if check(grid, x, i, n) {
				grid[x][i] = 1
				dfs(x + 1)
				grid[x][i] = 0
			}
		}
	}
	dfs(0)

	return ans
}

// 在grid[x][y] 上放着Q 会和之前有没有冲突
func check(grid [][]int, x, y, n int) bool {
	// 检查这一竖列
	for j := 0; j < y; j++ {
		if grid[x][j] > 0 {
			return false
		}
	}
	// 检查这一衡行
	for i := 0; i < x; i++ {
		if grid[i][y] > 0 {
			return false
		}
	}
	// 45度
	for i, j := x-1, y+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if grid[i][j] > 0 {
			return false
		}
	}
	// 135
	for i, j := x-1, y-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if grid[i][j] > 0 {
			return false
		}
	}

	return true
}
