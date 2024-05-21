package main

import (
	"sort"
)

func main() {

}

func pondSizes(land [][]int) []int {
	ans := make([]int, 0)
	for i, ch := range land {
		for j, ch2 := range ch {
			if ch2 != 0 {
				continue
			}

			ans = append(ans, dfs(land, i, j))
		}
	}
	sort.Ints(ans)
	if len(ans) > 0 && ans[0] == 0 {
		ans = ans[1:]
	}
	return ans
}

// 由垂直、水平或对角连接的水域为池塘
func dfs(grid [][]int, col, row int) int {
	if col < 0 || row < 0 || col >= len(grid) || row >= len(grid[col]) {
		return 0
	}
	if grid[col][row] != 0 {
		return 0
	}
	ans := 1
	grid[col][row] = 1
	ans += dfs(grid, col+1, row)
	ans += dfs(grid, col-1, row)
	ans += dfs(grid, col, row+1)
	ans += dfs(grid, col, row-1)
	// 对角线
	ans += dfs(grid, col+1, row+1)
	ans += dfs(grid, col-1, row-1)
	ans += dfs(grid, col+1, row-1)
	ans += dfs(grid, col-1, row+1)

	return ans
}
