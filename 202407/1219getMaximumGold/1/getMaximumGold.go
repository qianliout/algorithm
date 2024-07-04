package main

import "fmt"

func main() {
	fmt.Println(getMaximumGold([][]int{{0, 6, 0}, {5, 8, 7}, {0, 9, 0}}))
	fmt.Println(getMaximumGold([][]int{{1, 0, 7, 0, 0, 0}, {2, 0, 6, 0, 1, 0}, {3, 5, 6, 7, 4, 2}, {4, 3, 1, 0, 2, 0}, {3, 0, 5, 0, 20, 0}}))
}

// 这种写法是错的，但是不知道是什么原因出错的
func getMaximumGold(grid [][]int) int {
	var dfs func(visit [][]bool, i, j int) int

	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	dfs = func(visit [][]bool, i, j int) int {
		if !in(grid, i, j) {
			return 0
		}
		ans := grid[i][j]
		visit[i][j] = true
		for _, dir := range dirs {
			x, y := i+dir[0], j+dir[1]
			if !in(grid, x, y) {
				continue
			}
			if visit[x][y] {
				continue
			}
			// visit[x][y] = true
			ans = max(ans, grid[i][j]+dfs(visit, x, y))
		}
		return ans
	}
	mx := 0
	m, n := len(grid), len(grid[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] <= 0 {
				continue
			}
			visit := genVisit(m, n)
			res := dfs(visit, i, j)
			mx = max(mx, res)
		}
	}
	return mx
}

func genVisit(m, n int) [][]bool {
	visit := make([][]bool, m)
	for i := range visit {
		visit[i] = make([]bool, n)
	}
	return visit
}

func in(grid [][]int, x, y int) bool {
	m, n := len(grid), len(grid[0])
	if x < 0 || y < 0 {
		return false
	}
	if x >= m || y >= n {
		return false
	}
	if grid[x][y] <= 0 {
		return false
	}
	return true
}
