package main

func main() {

}

var dirs = [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

func containsCycle(grid [][]byte) bool {
	used := make([][]bool, len(grid))
	for i := range used {
		used[i] = make([]bool, len(grid[0]))
	}
	for i := range grid {
		for j := range grid[i] {
			ans := dfs(grid, i, j, i, j, 1, used)
			if ans {
				return true
			}

		}
	}
	return false
}

// 会超时
func dfs(grid [][]byte, sc, sr, col, row, cnt int, used [][]bool) bool {
	if col < 0 || row < 0 || col >= len(grid) || row >= len(grid[col]) {
		return false
	}

	used[col][row] = true
	ans := false
	for _, dir := range dirs {
		nc, nr := col+dir[0], row+dir[1]
		if !in(grid, nc, nr) {
			continue
		}
		if grid[nc][nr] != grid[sc][sr] {
			continue
		}
		if nc == sc && nr == sr {
			if cnt >= 4 {
				ans = true
				break
			}
		}
		if used[nc][nr] {
			continue
		}
		ans = ans || dfs(grid, sc, sr, nc, nr, cnt+1, used)
	}
	used[col][row] = false
	return ans
}

func in(grid [][]byte, col, row int) bool {
	if col < 0 || row < 0 || col >= len(grid) || row >= len(grid[col]) {
		return false
	}
	return true
}
