package main

func main() {

}

func numEnclaves(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	for c := 0; c < len(grid); c++ {
		dfs(grid, c, 0)
		dfs(grid, c, len(grid[0])-1)
	}
	for r := 0; r < len(grid[0]); r++ {
		dfs(grid, 0, r)
		dfs(grid, len(grid)-1, r)
	}
	ans := 0
	for _, ch := range grid {
		for _, ch2 := range ch {
			if ch2 == 1 {
				ans++
			}
		}
	}
	return ans
}

func dfs(grid [][]int, c, r int) {
	if c < 0 || r < 0 || c >= len(grid) || r >= len(grid[c]) {
		return
	}
	if grid[c][r] == 0 {
		return
	}
	grid[c][r] = 0
	dfs(grid, c+1, r)
	dfs(grid, c-1, r)
	dfs(grid, c, r+1)
	dfs(grid, c, r-1)
}
