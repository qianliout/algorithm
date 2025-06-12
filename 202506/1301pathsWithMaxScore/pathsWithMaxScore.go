package main

import (
	"fmt"
)

func main() {
	fmt.Println(pathsWithMaxScore([]string{"E12", "1X1", "21S"}))
}

func pathsWithMaxScore(board []string) []int {
	m, n := len(board), len(board[0])
	path := make([][]int, m+5)
	used := make([][]bool, m+5)
	for i := range path {
		path[i] = make([]int, n+5)
		used[i] = make([]bool, n+5)
	}
	//  dfs表示 从i,j 到0，0值获得的最大得分
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i < 0 || j < 0 || i >= m || j >= n || board[i][j] == 'S' {
			return 0
		}
		if used[i][j] {
			return 0
		}
		if path[i][j] > 0 {
			return path[i][j]
		}

		used[i][j] = true
		b := dfs(i-1, j)
		c := dfs(i-1, j-1)
		d := dfs(i, j-1)
		used[i][j] = false
		path[i][j] = int(board[i][j]) - int('0') + max(b, c, d)
		return path[i][j]

	}
	dfs(0, 0)
	mx := max(path[1][0], path[0][1])

	return []int{mx, 0}
}
