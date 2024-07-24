package main

import (
	"fmt"
)

func main() {
	fmt.Println(countSubIslands(
		[][]int{{1, 1, 1, 0, 0}, {0, 1, 1, 1, 1}, {0, 0, 0, 0, 0}, {1, 0, 0, 0, 0}, {1, 1, 0, 1, 1}},
		[][]int{{1, 1, 1, 0, 0}, {0, 0, 1, 1, 1}, {0, 1, 0, 0, 0}, {1, 0, 1, 1, 0}, {0, 1, 0, 1, 0}},
	))
}

// 它们只包含 0 （表示水域）和 1 （表示陆地
func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	m, n := len(grid1), len(grid1[0])
	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	// 从x,y点出发，能不能找到一个子岛屿
	visit := make([][]int, m)
	for i := range visit {
		visit[i] = make([]int, n)
	}
	var bfs func(x, y int) bool
	bfs = func(x, y int) bool {
		visit[x][y] = 1
		queue := [][]int{{x, y}}
		check := grid1[x][y] == 1
		for len(queue) > 0 {
			fir := queue[0]
			queue = queue[1:]
			a, b := fir[0], fir[1]
			for _, dir := range dirs {
				nx, ny := a+dir[0], b+dir[1]
				if nx < 0 || nx >= m || ny < 0 || ny >= n {
					continue
				}
				if grid2[nx][ny] == 0 {
					continue
				}
				if grid1[nx][ny] == 0 {
					// 这里不能直接返回，需要把这一个岛全部遍历完成
					check = false
				}
				if visit[nx][ny] == 1 {
					continue
				}
				visit[nx][ny] = 1
				queue = append(queue, []int{nx, ny})
			}
		}
		return check
	}
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid2[i][j] == 0 {
				continue
			}
			if visit[i][j] == 1 {
				continue
			}
			if bfs(i, j) {
				ans++
			}
		}
	}
	return ans
}
