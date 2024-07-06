package main

import "fmt"

func main() {
	fmt.Println(shiftGrid([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 12))
}

func shiftGrid(grid [][]int, k int) [][]int {
	m, n := len(grid), len(grid[0])
	ans := make([][]int, m)
	for i := range ans {
		ans[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// 这里一定要取余数，不然会越界
			x := ((i*n + j + k) / n) % m
			y := (i*n + j + k) % n
			ans[x][y] = grid[i][j]
		}
	}

	return ans
}
