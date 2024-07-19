package main

import (
	"sort"
)

func main() {

}

func largestSubmatrix(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	up := make([][]int, m)
	for i := range up {
		up[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 1 {
				if i == 0 {
					up[i][j] = 1
				} else {
					up[i][j] = up[i-1][j] + 1
				}
			}
		}
	}

	res := 0

	for i := 0; i < m; i++ {
		level := make([]int, n)
		for j := 0; j < n; j++ {
			level[j] = up[i][j]
		}
		sort.Ints(level)
		for j := 0; j < n; j++ {
			res = max(res, level[j]*(n-j))
		}
	}

	return res
}
