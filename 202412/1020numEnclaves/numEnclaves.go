package main

func main() {

}

func numEnclaves(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	var dfs func(i, j int)
	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	dfs = func(i, j int) {
		grid[i][j] = 0

		for _, dir := range dirs {
			x, y := i+dir[0], j+dir[1]
			if in(m, n, x, y) && grid[x][y] == 1 {
				dfs(x, y)
			}
		}
	}

	for i := 0; i < m; i++ {
		if grid[i][0] == 1 {
			dfs(i, 0)
		}
		if grid[i][n-1] == 1 {
			dfs(i, n-1)
		}
	}
	for j := 0; j < n; j++ {
		if grid[0][j] == 1 {
			dfs(0, j)
		}
		if grid[m-1][j] == 1 {
			dfs(m-1, j)
		}
	}
	cnt := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				cnt++
			}
		}
	}
	return cnt
}

func in(m, n, i, j int) bool {
	if i >= 0 && i < m && j >= 0 && j < n {
		return true
	}
	return false
}
