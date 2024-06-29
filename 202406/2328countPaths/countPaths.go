package main

import (
	"math"
)

func main() {

}

func countPaths(grid [][]int) int {
	var dfs func(i, j int) int
	mod := int(math.Pow10(9)) + 7
	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	m, n := len(grid), len(grid[0])

	mem := make([][]int, m)
	for i := range mem {
		mem[i] = make([]int, n)
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}
	dfs = func(i, j int) int {
		if !in(i, j, m, n) {
			return 0
		}
		if mem[i][j] != -1 {
			return mem[i][j]
		}

		res := 1 // 只有自已也算
		for _, dir := range dirs {
			x, y := i+dir[0], j+dir[1]
			// grid[x][y] > grid[i][j] 这里是严格大于，所以不会重复返问格子
			if in(x, y, m, n) && grid[x][y] > grid[i][j] {
				res = (res + dfs(x, y)) % mod
			}
		}
		mem[i][j] = res % mod
		return res % mod
	}
	sum := 0
	for i := range grid {
		for j := range grid[i] {
			sum = (sum + dfs(i, j)) % mod
		}
	}
	return sum
}

func in(i, j, m, n int) bool {
	if i < 0 || j < 0 {
		return false
	}
	if i >= m || j >= n {
		return false
	}
	return true
}
