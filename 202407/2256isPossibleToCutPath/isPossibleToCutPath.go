package main

func main() {

}

func isPossibleToCutPath(grid [][]int) bool {
	m, n := len(grid), len(grid[0])

	var dfs func(x, y int) bool

	dfs = func(x, y int) bool {
		if x == m-1 && y == n-1 {
			return true
		}
		// 把走过的下轮廓翻转，同时也避免走重复的路
		grid[x][y] = 0

		// 走下轮廓
		if x+1 < m && grid[x+1][y] == 1 {
			if dfs(x+1, y) {
				return true
			}
		}
		if y+1 < n && grid[x][y+1] == 1 {
			if dfs(x, y+1) {
				return true
			}
		}
		return false
	}
	if !dfs(0, 0) {
		// 本来都走不通，所以就直接不用翻转
		return true
	}

	// 上一次把下轮廓的值都翻转了，再走一次，如果 走不通，那说明有一个交集
	return !dfs(0, 0)
}
