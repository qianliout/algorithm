package main

import "strconv"

func main() {

}

func findColumnWidth(grid [][]int) []int {
	n := len(grid)
	m := len(grid[0])
	ans := make([]int, m)
	for i := 0; i < m; i++ {
		mx := 0
		for j := 0; j < n; j++ {
			mx = max(mx, len(strconv.Itoa(grid[j][i])))
		}
		ans[i] = mx
	}
	return ans
}
