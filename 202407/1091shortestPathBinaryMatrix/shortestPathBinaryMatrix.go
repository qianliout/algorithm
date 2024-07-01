package main

func main() {

}

func shortestPathBinaryMatrix(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	visit := make(map[pair]bool)
	queue := make([]pair, 0)
	if grid[0][0] == 0 {
		queue = append(queue, pair{0, 0})
		visit[pair{0, 0}] = true
	}
	ans := 0
	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {-1, -1}, {1, -1}, {-1, 1}, {1, 1}}

	for len(queue) > 0 {
		ans++
		lev := make([]pair, 0)
		for _, no := range queue {
			if no.x == m-1 && no.y == n-1 {
				return ans
			}

			for _, dir := range dirs {
				x, y := no.x+dir[0], no.y+dir[1]
				if in(m, n, x, y) && !visit[pair{x, y}] && grid[x][y] == 0 {
					p := pair{x, y}
					visit[p] = true
					lev = append(lev, p)
				}
			}
		}
		queue = lev
	}

	return -1
}

func in(m, n, x, y int) bool {
	if x < 0 || y < 0 {
		return false
	}
	if x >= m || y >= n {
		return false
	}
	return true
}

type pair struct {
	x, y int
}
