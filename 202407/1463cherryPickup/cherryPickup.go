package main

import (
	"fmt"
)

func main() {
	fmt.Println(cherryPickup([][]int{{3, 1, 1}, {2, 5, 1}, {1, 5, 5}, {2, 1, 1}}))
}

func cherryPickup(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	mem := make([][][]int, n)
	for i := range mem {
		mem[i] = make([][]int, m)
		for j := range mem[i] {
			mem[i][j] = make([]int, m)
			for k := range mem[i][j] {
				mem[i][j][k] = -1
			}
		}
	}

	// a:[i,j],b:[i,k]
	var dfs func(i, j, k int) int
	dfs = func(i, j, k int) int {
		if i < 0 || i >= n {
			return 0
		}
		if j < 0 || j >= m || k < 0 || k >= m {
			return 0
		}
		if mem[i][j][k] != -1 {
			return mem[i][j][k]
		}
		res := grid[i][j]
		if j != k {
			res += grid[i][k]
		}
		nex := 0
		for j0 := j - 1; j0 <= j+1; j0++ {
			for k0 := k - 1; k0 <= k+1; k0++ {
				nex = max(nex, dfs(i+1, j0, k0))
			}
		}
		mem[i][j][k] = res + nex
		return res + nex
	}
	res := dfs(0, 0, m-1)
	return res
}

// 错误写法
func cherryPickup2(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	// a:[i,j],b:[i,k]
	var dfs func(i, j, k int) int
	dirs := []int{-1, 0, 1}
	dfs = func(i, j, k int) int {
		if i <= 0 || i >= n {
			return 0
		}
		if j < 0 || j >= m || k < 0 || k >= m {
			return 0
		}
		res := grid[i][j]
		if j != k {
			res += grid[i][k]
		}

		// 这中写法是错误的
		nex := 0
		for _, dir := range dirs {
			y := j + dir
			nex = max(nex, dfs(i+1, y, k))
		}
		for _, dir := range dirs {
			y := k + dir
			nex = max(nex, dfs(i+1, j, y))
		}
		return res + nex
	}
	return dfs(0, 0, m-1)
}
