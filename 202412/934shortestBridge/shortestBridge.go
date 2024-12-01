package main

func main() {

}

func shortestBridge(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	ans := 0
	// 先把第一个岛变成2
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	var dfs func(i, j int)
	dfs = func(i, j int) {
		grid[i][j] = 2
		for _, dir := range dirs {
			x, y := i+dir[0], j+dir[1]
			if in(m, n, x, y) && grid[x][y] == 1 {
				dfs(x, y)
			}
		}
	}
	for i := 0; i < m; i++ {
		find1 := false
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				dfs(i, j)
				find1 = true
				break
			}
		}
		if find1 {
			break
		}
	}
	q := make([]pair, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				q = append(q, pair{i, j, 0})
			}
		}
	}
	for len(q) > 0 {
		ans++
		lev := make([]pair, 0)
		for _, no := range q {
			for _, dir := range dirs {
				x, y := no.x+dir[0], no.y+dir[1]
				if in(m, n, x, y) {
					if grid[x][y] == 2 {
						return no.dis
					}
					if grid[x][y] == 0 {
						grid[x][y] = 1 // 防止重复访问
						lev = append(lev, pair{x, y, no.dis + 1})
					}
				}
			}
		}
		q = lev
	}
	return -1
}

func in(m, n, i, j int) bool {
	if i >= 0 && i < m && j >= 0 && j < n {
		return true
	}
	return false
}

type pair struct {
	x, y int
	dis  int
}
