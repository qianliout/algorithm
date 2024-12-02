package main

func main() {

}

func maxAreaOfIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dirs := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if !in(m, n, i, j) || grid[i][j] == 0 {
			return 0
		}
		ans := 1
		grid[i][j] = 0
		for _, dir := range dirs {
			ans += dfs(i+dir[0], j+dir[1])
		}
		return ans
	}
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				ans = max(ans, dfs(i, j))
			}
		}
	}
	return ans
}
func in(m, n, i, j int) bool {
	if i >= 0 && i < m && j >= 0 && j < n {
		return true
	}
	return false
}

//  1 (代表土地)
