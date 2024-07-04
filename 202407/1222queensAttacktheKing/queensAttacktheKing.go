package main

func main() {

}

func queensAttacktheKing(queens [][]int, king []int) [][]int {
	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {-1, -1}, {1, -1}, {-1, 1}, {1, 1}}
	ans := make([][]int, 0)

	g := make([][]int, 8)
	for i := range g {
		g[i] = make([]int, 8)
	}
	for _, ch := range queens {
		g[ch[0]][ch[1]] = 1
	}

	for _, dir := range dirs {
		dep := 1
		for {
			dx, dy := dir[0]*dep, dir[1]*dep
			x, y := king[0]+dx, king[1]+dy
			if !in(g, x, y) {
				break
			}
			if g[x][y] == 1 {
				ans = append(ans, []int{x, y})
				// 找到就不能再找了，挡住了
				break
			}
			dep++
		}
	}
	return ans
}

func in(grid [][]int, x, y int) bool {
	m, n := len(grid), len(grid[0])
	if x < 0 || y < 0 {
		return false
	}
	if x >= m || y >= n {
		return false
	}
	return true
}
