package main

func main() {

}

func maxAreaOfIsland(grid [][]int) int {
	ans := 0
	for i, ch := range grid {
		for j := range ch {
			ans = max(ans, dfs(grid, i, j))
		}
	}
	return ans
}

func dfs(grid [][]int, col, row int) int {
	if col < 0 || row < 0 || col >= len(grid) || row >= len(grid[col]) {
		return 0
	}
	if grid[col][row] != 1 {
		return 0
	}
	ans := 1
	grid[col][row] = 0
	ans += dfs(grid, col+1, row)
	ans += dfs(grid, col-1, row)
	ans += dfs(grid, col, row+1)
	ans += dfs(grid, col, row-1)
	return ans
}
