package main

import "fmt"

func main() {
	fmt.Println(countServers([][]int{{1, 0}, {1, 1}}))
}

func countServers(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	col, row := make([]int, m), make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				col[i]++
				row[j]++
			}

		}
	}
	ans := 0
	for i := 0; i < m; i++ {
		for j, ch := range grid[i] {
			if ch == 1 {
				if col[i] > 1 || row[j] > 1 {
					ans++
				}
			}
		}
	}
	return ans
}
