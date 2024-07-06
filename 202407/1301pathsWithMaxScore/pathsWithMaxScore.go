package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(pathsWithMaxScore([]string{"E23", "2X2", "12S"}))
}

func pathsWithMaxScore(board []string) []int {
	mod := int(math.Pow10(9)) + 7
	m := len(board)
	grid := make([][]byte, m)
	for i := range board {
		grid[i] = []byte(board[i])
	}
	n := len(board[0])
	var dfs func(i, j int) int
	mem := make([][]int, m)
	for i := range mem {
		mem[i] = make([]int, n)
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}

	dfs = func(i, j int) int {
		if !in(m, n, i, j) {
			return 0
		}
		if grid[i][j] == 'X' {
			return 0
		}
		if mem[i][j] != -1 {
			return mem[i][j]
		}
		num := int(grid[i][j] - '0')
		res := 0

		if in(m, n, i+1, j) && grid[i+1][j] != 'X' {
			res = max(res, dfs(i+1, j)+num)
		}
		if in(m, n, i, j+1) && grid[i][j+1] != 'X' {
			res = max(res, dfs(i, j+1)+num)
		}
		if in(m, n, i+1, j+1) && grid[i+1][j+1] != 'X' {
			res = max(res, dfs(i+1, j+1)+num)
		}
		mem[i][j] = res % mod
		return res % mod
	}

	mem2 := make([][]int, m)
	for i := range mem2 {
		mem2[i] = make([]int, n)
		for j := range mem2[i] {
			mem2[i][j] = -1
		}
	}

	var path func(i, j int) int

	path = func(i, j int) int {
		if i == 0 && j == 0 {
			return 1
		}
		if !in(m, n, i, j) {
			return 0
		}
		if grid[i][j] == 'X' {
			return 0
		}
		if mem2[i][j] != -1 {
			return mem2[i][j]
		}
		num := int(grid[i][j] - '0')

		res := 0

		if in(m, n, i+1, j) && grid[i+1][j] != 'X' && dfs(i+1, j)+num == dfs(i, j) {
			res++
		}
		if in(m, n, i, j+1) && grid[i][j+1] != 'X' && dfs(i, j+1)+num == dfs(i, j) {
			res++
		}
		if in(m, n, i+1, j+1) && grid[i+1][j+1] != 'X' && dfs(i+1, j+1)+num == dfs(i, j) {
			res++
		}
		mem2[i][j] = res
		return res
	}
	return []int{dfs(0, 0), path(0, 0)}
}

func in(m, n, i, j int) bool {
	if i < 0 || j < 0 {
		return false
	}
	if i >= m || j >= n {
		return false
	}
	return true
}
