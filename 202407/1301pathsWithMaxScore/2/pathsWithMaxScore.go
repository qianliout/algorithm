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
	dist := make([][]int, m)
	for i := range dist {
		dist[i] = make([]int, n)
	}

	dirs := [][]int{{1, 0}, {0, 1}, {1, 1}}
	// 求出每个点的最大分数
	var dfs1 func(i, j, s int)
	dfs1 = func(i, j, s int) {
		if !in(m, n, i, j) {
			return
		}
		if i == m-1 && j == m-1 {
			dist[i][j] = max(dist[i][j], s)
			return
		}
		tem := 0
		if i == 0 && j == 0 {
			tem = 0
		} else if grid[i][j] == 'X' {
			// 按递归的逻辑是不会有这种情况的
			return
		} else {
			tem = int(grid[i][j] - '0')
		}
		if dist[i][j] >= s+tem && dist[i][j] > 0 {
			return // 剪枝
		}

		dist[i][j] = s + tem
		for _, d := range dirs {
			x, y := i+d[0], j+d[1]
			if in(m, n, x, y) && grid[x][y] != 'X' {
				dfs1(x, y, s+tem)
			}
		}
	}
	// 找得分最少，从最后向前找,和题意一样
	dfs1(0, 0, 0)

	var dfs2 func(i, j int) int
	path := make([][]int, m)
	for i := range path {
		path[i] = make([]int, n)
		for j := range path[i] {
			path[i][j] = -1
		}
	}
	dfs2 = func(i, j int) int {
		if i == m-1 && j == n-1 {
			return 1
		}
		if !in(m, n, i, j) {
			return 0
		}
		if path[i][j] != -1 {
			return path[i][j]
		}
		res := 0
		for _, d := range dirs {
			x, y := i+d[0], j+d[1]
			if in(m, n, x, y) && grid[x][y] != 'X' {
				nex := 0
				if grid[x][y] != 'S' {
					nex = int(grid[x][y] - '0')
				}
				if dist[i][j]+nex == dist[x][y] {
					res += dfs2(x, y)
				}
			}
		}
		path[i][j] = res
		return res
	}
	pathCnt := dfs2(0, 0)

	return []int{dist[m-1][n-1], pathCnt % mod}
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
