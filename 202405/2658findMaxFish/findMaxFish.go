package main

func main() {

}

func findMaxFish(grid [][]int) int {
	ans := 0
	for i, ch := range grid {
		for j, ch1 := range ch {
			if ch1 == 0 {
				continue
			}
			ans = max(ans, dfs(grid, i, j))
		}
	}
	return ans
}

func dfs(grid [][]int, col, row int) int {
	if col < 0 || row < 0 || col >= len(grid) || row >= len(grid[col]) {
		return 0
	}
	if grid[col][row] == 0 {
		return 0
	}
	ans := grid[col][row]
	grid[col][row] = 0
	ans += dfs(grid, col+1, row)
	ans += dfs(grid, col-1, row)
	ans += dfs(grid, col, row+1)
	ans += dfs(grid, col, row-1)
	return ans
}
