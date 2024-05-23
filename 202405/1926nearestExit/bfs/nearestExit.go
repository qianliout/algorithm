package main

import (
	"fmt"
)

func main() {
	// maze := [][]byte{[]byte("+++"), []byte("..."), []byte("+++")}
	maze := [][]byte{[]byte("..")}

	fmt.Println(nearestExit(maze, []int{0, 1}))
}

var dirs = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func nearestExit(maze [][]byte, entrance []int) int {
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
	ans := 0
	mem[entrance[0]][entrance[1]] = 1

	for len(queue) > 0 {
		ans++
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
					return ans
				}

				if maze[nx][ny] == '+' {
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

func in(grid [][]byte, col, row int) bool {
	if col < 0 || row < 0 || col >= len(grid) || row >= len(grid[col]) {
		return false
	}
	return true
}

func exit(grid [][]byte, col, row int) bool {
	if !in(grid, col, row) {
		return false
	}
	if grid[col][row] == '+' {
		return false
	}
	if col == 0 || col == len(grid)-1 || row == 0 || row == len(grid[col])-1 {
		return true
	}

	return false
}
