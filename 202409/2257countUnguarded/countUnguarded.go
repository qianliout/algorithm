package main

func main() {

}

func countUnguarded(m int, n int, guards [][]int, walls [][]int) int {
	grid := make([][]byte, m)
	for i := range grid {
		grid[i] = make([]byte, n)
	}
	for _, ch := range guards {
		grid[ch[0]][ch[1]] = 'G'
	}
	for _, ch := range walls {
		grid[ch[0]][ch[1]] = 'W'
	}
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	for _, ch := range guards {
		x, y := ch[0], ch[1]
		for _, dir := range dirs {
			nx := x + dir[0]
			ny := y + dir[1]
			for nx >= 0 && nx < m && ny >= 0 && ny < n && grid[nx][ny] != 'G' && grid[nx][ny] != 'W' {
				grid[nx][ny] = 'I'
				nx = nx + dir[0]
				ny = ny + dir[1]
			}
		}
	}
	ans := 0
	for _, row := range grid {
		for _, col := range row {
			if col == byte(0) {
				ans++
			}
		}
	}
	return ans
}
