package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(pathsWithMaxScore([]string{"E23", "2X2", "12S"}))
	fmt.Println(pathsWithMaxScore([]string{"E12", "1X1", "21S"}))
}

func pathsWithMaxScore(board []string) []int {
	mod := int(math.Pow10(9)) + 7
	m := len(board)
	grid := make([][]byte, m)
	for i := range board {
		grid[i] = []byte(board[i])
	}
	n := len(board[0])
	var dfs func(i, j, s int)
	mem := make([][]int, m)
	for i := range mem {
		mem[i] = make([]int, n)
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}

	dfs = func(i, j, s int) {
		if !in(m, n, i, j) {
			return
		}
		if grid[i][j] == 'X' {
			return
		}
		mem[i][j] = max(mem[i][j], s)
		a := 0
		if grid[i][j] >= '0' && grid[i][j] <= '9' {
			a = int(grid[i][j] - '0')
		}
		dfs(i-1, j, s+a)
		dfs(i, j-1, s+a)
		dfs(i-1, j-1, s+a)
	}
	dfs(m-1, n-1, 0)

	var path func(i, j int)
	pathCnt := 0
	path = func(i, j int) {
		if !in(m, n, i, j) {
			return
		}
		a := 0
		if grid[i][j] >= '0' && grid[i][j] <= '9' {
			a = int(grid[i][j] - '0')
		}
		s := mem[i][j]
		if in(m, n, i-1, j) && mem[i-1][j]+a == s {
			path(i-1, j)
			pathCnt++
		}
		if in(m, n, i, j-1) && mem[i][j-1]+a == s {
			pathCnt++
			path(i, j-1)
		}
		if in(m, n, i-1, j-1) && mem[i-1][j-1]+a == s {
			pathCnt++
			path(i, j-1)
		}
	}
	path(m-1, n-1)

	return []int{mem[0][0] % mod, pathCnt / 2}
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
