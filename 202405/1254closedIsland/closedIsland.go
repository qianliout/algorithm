package main

func main() {

}

func closedIsland(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])

	// 先从边界出发,把水全部改成土地
	for c := 0; c < m; c++ {
		dfs(grid, c, 0)
		dfs(grid, c, n-1)
	}
	for r := 0; r < len(grid[0]); r++ {
		dfs(grid, 0, r)
		dfs(grid, m-1, r)
	}

	ans := 0

	for i, ch := range grid {
		for j, ch2 := range ch {
			if ch2 == 0 {
				ans++
				dfs(grid, i, j)
			}
		}
	}
	return ans
}

// 0 是土地，1是水
func dfs(grid [][]int, c, r int) {
	if c < 0 || r < 0 || c >= len(grid) || r >= len(grid[c]) {
		return
	}
	if grid[c][r] != 0 {
		return
	}
	grid[c][r] = 1 // 和边

	dfs(grid, c+1, r)
	dfs(grid, c-1, r)
	dfs(grid, c, r+1)
	dfs(grid, c, r-1)
}
