package main

import (
	"fmt"
)

func main() {
	grid := [][]int{{2, 4, 3, 5}, {5, 4, 9, 3}, {3, 4, 2, 11}, {10, 9, 13, 15}}
	fmt.Println(maxMoves(grid))
}

var dirs = [][]int{{-1, 1}, {0, 1}, {1, 1}}

func maxMoves(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	ans := 1
	mem := make([][]int, len(grid))
	for i := range mem {
		mem[i] = make([]int, len(grid[0]))
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}
	// 你可以从矩阵第一列中的 任一 单元格出发
	for i := range grid {
		ans = max(ans, dfs(grid, i, 0, mem))
	}
	return ans - 1
}

// 不加 mem 会超时
func dfs(grid [][]int, col, row int, mem [][]int) int {
	if mem[col][row] != -1 {
		return mem[col][row]
	}
	ans := 1
	for _, dir := range dirs {
		x, y := col+dir[0], row+dir[1]
		if !in(grid, x, y) {
			continue
		}
		if grid[x][y] > grid[col][row] {
			ans = max(ans, dfs(grid, x, y, mem)+1)
		}
	}
	mem[col][row] = ans
	return ans
}

func in(grid [][]int, col, row int) bool {
	if col < 0 || row < 0 || col >= len(grid) || row >= len(grid[col]) {
		return false
	}
	return true
}
