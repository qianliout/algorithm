package main

func main() {

}

var dirs = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func maxDistance(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m := len(grid)
	queue := make([][]int, 0)
	mem := make(map[int]int)
	for i := range grid {
		for j, ch := range grid[i] {
			if ch == 1 {
				queue = append(queue, []int{i, j})
				mem[i*m+j] = 0
			}
		}
	}
	ans := -1
	for len(queue) > 0 {
		first := queue[0]
		queue = queue[1:]
		x, y := first[0], first[1]
		step := mem[x*m+y]
		for _, dir := range dirs {
			nx, ny := x+dir[0], y+dir[1]
			if !in(grid, nx, ny) {
				continue
			}
			// 说明已经来过了
			if grid[nx][ny] != 0 {
				continue
			}
			grid[nx][ny] = step + 1
			queue = append(queue, []int{nx, ny})
			mem[nx*m+ny] = step + 1
			ans = max(ans, step+1)
		}
	}
	return ans

}

func in(grid [][]int, col, row int) bool {
	if col < 0 || row < 0 || col >= len(grid) || row >= len(grid[col]) {
		return false
	}
	return true
}
