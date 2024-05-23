package main

import (
	"fmt"

	. "outback/algorithm/common/unionfind"
)

func main() {
	// grid1 := [][]int{{1, 0}, {0, 1}}
	grid2 := [][]int{{1, 1, 1}, {0, 0, 0}, {1, 1, 1}}
	// grid3 := [][]int{{1, 1}, {1, 1}}
	grid4 := [][]int{{0, 0, 0, 0, 0, 0, 0}, {0, 1, 1, 1, 1, 0, 0}, {0, 1, 0, 0, 1, 0, 0},
		{1, 0, 1, 0, 1, 0, 0}, {0, 1, 0, 0, 1, 0, 0}, {0, 1, 0, 0, 1, 0, 0}, {0, 1, 1, 1, 1, 0, 0}}

	// fmt.Println(largestIsland(grid1))
	fmt.Println(largestIsland(grid2))
	// fmt.Println(largestIsland(grid3))
	fmt.Println(largestIsland(grid4))
}

func largestIsland(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	dirs := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	uf := NewSizeUnionFind(m*n + 1)

	for i := range grid {
		for j, ch := range grid[i] {
			if ch == 0 {
				continue
			}
			for _, dir := range dirs {
				x, y := i+dir[0], j+dir[1]
				if !in(grid, x, y) || grid[x][y] == 0 {
					continue
				}
				if uf.IsConnected(i*m+j, x*m+y) {
					continue
				}
				uf.Union(i*m+j, x*m+y)
			}
		}
	}

	ans := 0
	for i := range grid {
		for j, ch := range grid[i] {
			if ch == 1 {
				ans = max(ans, uf.Size[uf.Find(i*m+j)])
				continue
			}
			tot := 1
			used := make(map[int]bool)
			for _, dir := range dirs {
				x, y := i+dir[0], j+dir[1]
				if !in(grid, x, y) || grid[x][y] == 0 {
					continue
				}
				root := uf.Find(x*m + y)
				if used[root] {
					continue
				}
				tot += uf.Size[uf.Find(root)]

				used[root] = true
			}
			ans = max(ans, tot)
		}
	}
	return ans
}

func in(grid [][]int, col, row int) bool {
	if col < 0 || row < 0 || col >= len(grid) || row >= len(grid[col]) {
		return false
	}
	return true
}
