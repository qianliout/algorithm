package main

func main() {

}

func oddCells(m int, n int, indices [][]int) int {
	grid := make([][]int, m)
	for i := range grid {
		grid[i] = make([]int, n)
	}
	for _, ch := range indices {
		x, y := ch[0], ch[1]
		for i := 0; i < n; i++ {
			grid[x][i]++
		}
		for i := 0; i < m; i++ {
			grid[i][y]++
		}
	}
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j]&1 == 1 {
				ans++
			}
		}
	}

	return ans
}
