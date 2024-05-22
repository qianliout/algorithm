package main

import (
	"fmt"
)

func main() {
	fmt.Println(hasValidPath([][]int{{2, 4, 3}, {6, 5, 2}}))
	fmt.Println(hasValidPath([][]int{{4, 3, 3}, {6, 5, 2}}))
}

var dx = [4]int{1, 0, -1, 0}
var dy = [4]int{0, 1, 0, -1}

var pipe = [7][4]int{{-1, -1, -1, -1}, {-1, 1, -1, 3}, {0, -1, 2, -1}, {-1, 0, 3, -1}, {-1, -1, 1, 0}, {3, 2, -1, -1}, {1, -1, -1, 2}}

func hasValidPath(grid [][]int) bool {
	sta := grid[0][0]
	for i := 0; i < 4; i++ {
		if pipe[sta][i] != -1 {
			vist := make([][]bool, len(grid))
			for i := range vist {
				vist[i] = make([]bool, len(grid[0]))
			}

			if dfs(grid, 0, 0, pipe[sta][i], vist) {
				return true
			}
		}
	}
	return false
}

func dfs(grid [][]int, col, row int, dir int, vist [][]bool) bool {
	if vist[col][row] {
		return false
	}
	if col < 0 || row < 0 || col >= len(grid) || row >= len(grid[0]) {
		return false
	}

	if col == len(grid)-1 && row == len(grid[0])-1 {
		return true
	}
	nc, nr := col+dx[dir], row+dy[dir]
	if nc < 0 || nr < 0 || nc >= len(grid) || nr >= len(grid[0]) {
		return false
	}
	vist[col][row] = true

	nex := grid[nc][nr]
	if pipe[nex][dir] != -1 {
		return dfs(grid, nc, nr, pipe[nex][dir], vist)
	}
	return false
}
