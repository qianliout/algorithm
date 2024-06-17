package main

import (
	"slices"
)

func main() {

}

func luckyNumbers(matrix [][]int) []int {
	n, m := len(matrix), len(matrix[0])
	col := make([]int, n)
	row := make([]int, m)
	for i := 0; i < n; i++ {
		col[i] = slices.Min(matrix[i])
	}
	for i := 0; i < m; i++ {
		ans := 0
		for j := 0; j < n; j++ {
			ans = max(ans, matrix[j][i])
		}
		row[i] = ans
	}
	ans := make([]int, 0)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			d := matrix[i][j]
			if d == col[i] && d == row[j] {
				ans = append(ans, d)
			}
		}
	}
	return ans
}
