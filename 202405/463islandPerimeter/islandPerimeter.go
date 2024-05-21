package main

func main() {

}

// 中间没有湖，使用问题简单的多
// 对于一个陆地格子的每条边，它被算作岛屿的周长当且仅当这条边为网格的边界或者相邻的另一个格子为水域。 因此，我们可以遍历每个陆地格子，
// 看其四个方向是否为边界或者水域，如果是，将这条边的贡献（即1）加入答案 ans 中即可。
// 也就是说从一个陆地格子出发，如果走到边边界上，说明有个条了，如果走到水域了，说明也有个边了
func islandPerimeter(grid [][]int) int {
	ans := 0
	for i, ch := range grid {
		for j, ch2 := range ch {
			if ch2 != 1 {
				continue
			}
			ans += dfs(grid, i, j)
		}
	}
	return ans
}

func dfs(grid [][]int, col, row int) int {
	if col < 0 || row < 0 || col >= len(grid) || row >= len(grid[col]) || grid[col][row] == 0 {
		return 1
	}

	if grid[col][row] == 2 {
		return 0
	}
	ans := 1
	grid[col][row] = 2
	ans += dfs(grid, col+1, row)
	ans += dfs(grid, col-1, row)
	ans += dfs(grid, col, row+1)
	ans += dfs(grid, col, row-1)
	return ans
}
