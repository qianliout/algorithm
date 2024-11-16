package main

func main() {

}

func numIslands(grid [][]byte) int {
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	m, n := len(grid), len(grid[0])
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if grid[i][j] == '0' {
			return
		}
		grid[i][j] = '0'
		for _, dir := range dirs {
			x, y := i+dir[0], j+dir[1]
			if x < 0 || y < 0 || x >= m || y >= n {
				continue
			}
			dfs(x, y)
		}
	}
	ans := 0
	for i := range grid {
		for j, ch := range grid[i] {
			if ch == '1' {
				ans++
				dfs(i, j)
			}
		}
	}
	return ans
}
