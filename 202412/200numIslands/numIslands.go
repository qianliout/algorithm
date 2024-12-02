package main

func main() {

}

func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	var dfs func(i, j int)
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	dfs = func(i, j int) {
		grid[i][j] = '0'
		for _, dir := range dirs {
			x, y := i+dir[0], j+dir[1]
			if !in(m, n, x, y) || grid[x][y] == '0' {
				continue
			}
			dfs(x, y)
		}
	}
	cnt := 0
	for i := range grid {
		for j, ch := range grid[i] {
			if ch == '1' {
				cnt++
				dfs(i, j)
			}
		}
	}
	return cnt
}

func in(m, n, i, j int) bool {
	if i < 0 || i >= m || j < 0 || j >= n {
		return false
	}
	return true
}
