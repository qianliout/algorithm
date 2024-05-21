package main

import (
	"fmt"
)

func main() {
	grid := [][]int{{2, 2, 2}, {2, 2, 2}, {2, 2, 2}}
	ans := colorBorder(grid, 1, 1, 3)
	fmt.Println(ans)
}

func colorBorder(grid [][]int, row int, col int, color int) [][]int {
	if col < 0 || row < 0 || row >= len(grid) || col >= len(grid[row]) {
		return grid
	}
	used := make([][]int, len(grid))
	for i := range used {
		used[i] = make([]int, len(grid[i]))
	}

	dfs(grid, row, col, color, used)
	for i, ch := range used {
		for j, ch2 := range ch {
			if ch2 == 0 {
				used[i][j] = grid[i][j]
			}
		}
	}
	return used
}

func dfs(grid [][]int, col, row int, color int, used [][]int) {
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
		if used[nc][nr] != 0 {
			// 已经遍历过了
			continue
		}
		used[nc][nr] = -1
		dfs(grid, nc, nr, color, used)

	}
	if cnt != 4 {
		used[col][row] = color
	} else {
		used[col][row] = grid[col][row]
	}
}
