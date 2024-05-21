package main

import (
	"fmt"
)

func main() {
	grid := [][]int{{2, 2, 2}, {2, 2, 2}, {2, 2, 2}}
	ans := colorBorder(grid, 1, 1, 2)
	fmt.Println(ans)

}

func colorBorder(grid [][]int, row int, col int, color int) [][]int {
	if col < 0 || row < 0 || row >= len(grid) || col >= len(grid[row]) {
		return grid
	}
	pre := grid[row][col]
	used := make([][]int, len(grid))
	for i := range used {
		used[i] = make([]int, len(grid[i]))
	}

	dfs(grid, row, col, pre, color, used)
	return grid
}

// 连通分量的边界 进行着色。在内部的话是不用着色的
func dfs(grid [][]int, col, row int, pre int, color int, used [][]int) {

	if col < 0 || row < 0 || col >= len(grid) || row >= len(grid[col]) {
		return
	}
	if used[col][row] > 0 {
		return
	}
	used[col][row] = 1

	if grid[col][row] != pre {
		return
	}

	if !check(grid, col, row) {
		grid[col][row] = color
	}

	dfs(grid, col+1, row, pre, color, used)
	dfs(grid, col-1, row, pre, color, used)
	dfs(grid, col, row+1, pre, color, used)
	dfs(grid, col, row-1, pre, color, used)
}

func check(grid [][]int, col, row int) bool {
	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	cnt := 0
	for _, dir := range dirs {
		nc, nr := col+dir[0], row+dir[1]
		if nc < 0 || nr < 0 || nc >= len(grid) || nr >= len(grid[nc]) {
			continue
		}
		if grid[nc][nr] != grid[col][row] {
			continue
		}
		cnt++
	}
	return cnt != 4
}
