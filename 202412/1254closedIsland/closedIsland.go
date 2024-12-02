package main

func main() {

}

// 0是土地，1是水

func closedIsland(grid [][]int) int {
	full(grid)
	cnt := 0
	m, n := len(grid), len(grid[0])
	var dfs func(i, j int)
	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	dfs = func(i, j int) {
		grid[i][j] = 1

		for _, dir := range dirs {
			x, y := i+dir[0], j+dir[1]
			if in(m, n, x, y) && grid[x][y] == 0 {
				dfs(x, y)
			}
		}
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 {
				cnt++
				dfs(i, j)
			}
		}
	}
	return cnt

}

func full(grid [][]int) {
	m, n := len(grid), len(grid[0])
	var dfs func(i, j int)
	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	dfs = func(i, j int) {
		grid[i][j] = 1

		for _, dir := range dirs {
			x, y := i+dir[0], j+dir[1]
			if in(m, n, x, y) && grid[x][y] == 0 {
				dfs(x, y)
			}
		}
	}

	for i := 0; i < m; i++ {
		if grid[i][0] == 0 {
			dfs(i, 0)
		}
		if grid[i][n-1] == 0 {
			dfs(i, n-1)
		}
	}
	for j := 0; j < n; j++ {
		if grid[0][j] == 0 {
			dfs(0, j)
		}
		if grid[m-1][j] == 0 {
			dfs(m-1, j)
		}
	}
}

func in(m, n, i, j int) bool {
	if i >= 0 && i < m && j >= 0 && j < n {
		return true
	}
	return false
}
