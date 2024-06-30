package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minimumObstacles([][]int{{0, 1, 1}, {1, 1, 0}, {1, 1, 0}}))
	fmt.Println(minimumObstacles([][]int{{0, 1, 0, 0, 0}, {0, 1, 0, 1, 0}, {0, 0, 0, 1, 0}}))
}

// 也不能获取到正确的答案
func minimumObstacles(grid [][]int) int {
	inf := math.MaxInt / 100
	m, n := len(grid), len(grid[0])
	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	var dfs func(i, j int) int
	visit := make([][]int, m)
	for i := range visit {
		visit[i] = make([]int, n)
	}

	dfs = func(i, j int) int {
		if !in(m, n, i, j) {
			return inf
		}
		visit[i][j] = 1

		if i == 0 && j == 0 {
			return grid[0][0]
		}
		res := inf
		for _, dir := range dirs {
			x, y := i+dir[0], j+dir[1]
			if in(m, n, x, y) && visit[x][y] == 0 {
				visit[x][y] = 1
				res = min(res, dfs(x, y))
			}
		}
		if grid[i][j] == 1 {
			res++
		}
		return res
	}

	return dfs(m-1, n-1)
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
