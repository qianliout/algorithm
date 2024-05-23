package main

import (
	. "outback/algorithm/common/utils"
)

func main() {

}

var dirs = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func maxDistance(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	ans := -1
	for i := range grid {
		for j, ch := range grid[i] {
			if ch == 1 {
				continue
			}
			ans = max(ans, nearestExit(grid, []int{i, j}))
		}
	}
	return ans
}

func buidMem(grid [][]int) [][]int {
	mem := make([][]int, len(grid))
	for i := range mem {
		mem[i] = make([]int, len(grid[0]))
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}
	return mem
}

func nearestExit(maze [][]int, entrance []int) int {
	if len(maze) == 0 || len(maze[0]) == 0 {
		return 0
	}

	mem := make([][]int, len(maze))
	for i := range mem {
		mem[i] = make([]int, len(maze[0]))
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}

	queue := make([][]int, 0)
	queue = append(queue, entrance)
	mem[entrance[0]][entrance[1]] = 1

	for len(queue) > 0 {
		lev := make([][]int, 0)
		for _, no := range queue {
			x, y := no[0], no[1]

			for _, dir := range dirs {
				nx, ny := x+dir[0], y+dir[1]
				if !in(maze, nx, ny) {
					continue
				}
				// 这一步是防止从原来点出去
				if mem[nx][ny] == 1 {
					continue
				}
				if exit(maze, nx, ny) {
					dis := Abs(nx-entrance[0]) + Abs(ny-entrance[1])
					return dis
				}

				if maze[nx][ny] == 1 {
					continue
				}

				lev = append(lev, []int{nx, ny})
				mem[nx][ny] = 1
			}
		}

		queue = lev
	}
	return -1
}

func in(grid [][]int, col, row int) bool {
	if col < 0 || row < 0 || col >= len(grid) || row >= len(grid[col]) {
		return false
	}
	return true
}

func exit(grid [][]int, col, row int) bool {
	if !in(grid, col, row) {
		return false
	}
	if grid[col][row] == 0 {
		return false
	}
	return true
}
