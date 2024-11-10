package main

import (
	"fmt"
)

func main() {
	fmt.Println(numberOfRightTriangles([][]int{{0}, {0}}))

}

func numberOfRightTriangles(grid [][]int) int64 {
	m, n := len(grid), len(grid[0])
	// hen
	row := make([]int, m)
	for i := range grid {
		for j := range grid[i] {
			row[i] += grid[i][j]
		}
	}
	// shu
	col := make([]int, n)
	for j := 0; j < n; j++ {
		for i := 0; i < m; i++ {
			col[j] += grid[i][j]
		}
	}
	ans := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				ans += max(0, row[i]-1) * max(0, col[j]-1)
			}
		}
	}
	return int64(ans)
}
