package main

import (
	"fmt"
)

func main() {
	fmt.Println(shortestBridge([][]int{{0, 1, 0}, {0, 0, 0}, {0, 0, 1}}))
}

func shortestBridge(grid [][]int) int {
	dirs := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

	n, m := len(grid), len(grid[0])

	for i := 0; i < n; i++ {
		find := false
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				find = true
				dfs(grid, i, j) // 说了只有两个桥，就先执行第一个桥
				break
			}
		}
		if find {
			break
		}
	}
	// 再执行 bfs
	ans := 0
	queue := make([]pair, 0)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				queue = append(queue, pair{i, j})
			}
		}
	}
	for len(queue) > 0 {
		lev := make([]pair, 0)
		for _, no := range queue {
			x, y := no.col, no.row
			for _, dir := range dirs {
				nx, ny := x+dir[0], y+dir[1]
				if nx < 0 || ny < 0 || nx >= n || ny >= m {
					continue
				}
				if grid[nx][ny] == 2 {
					return ans
				}
				if grid[nx][ny] == 0 {
					grid[nx][ny] = 1
					lev = append(lev, pair{nx, ny})
				}
			}
		}
		if len(lev) > 0 {
			ans++
		}
		queue = lev
	}
	return ans
}

type pair struct {
	col, row int
}

func dfs(grid [][]int, i, j int) {
	n, m := len(grid), len(grid[0])
	if i < 0 || j < 0 || i >= n || j >= m {
		return
	}
	if grid[i][j] == 0 {
		return
	}
	if grid[i][j] == 1 {
		grid[i][j] = 2
		dfs(grid, i+1, j)
		dfs(grid, i-1, j)
		dfs(grid, i, j+1)
		dfs(grid, i, j-1)
	}
}

/*
输入：s = "IDID"
输出：[0,4,1,3,2]
*/
func diStringMatch(s string) []int {
	n := len(s)
	le, ri := 0, n
	ans := make([]int, n+1)
	for i := 0; i < n; i++ {
		if s[i] == 'I' {
			ans[i] = le
			le++
		} else if s[i] == 'D' {
			ans[i] = ri
			ri--
		}
	}
	ans[n] = le
	return ans
}
