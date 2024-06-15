package main

import (
	"fmt"
)

func main() {
	fmt.Println(countSquares([][]int{{0, 1, 1, 1}, {1, 1, 1, 1}, {0, 1, 1, 1}}))
}

func countSquares(matrix [][]int) int {
	n, m := len(matrix), len(matrix[0])
	pre := make([][]int, n+1)
	for i := range pre {
		pre[i] = make([]int, m+1)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			pre[i+1][j+1] = pre[i+1][j] + pre[i][j+1] - pre[i][j] + matrix[i][j]
		}
	}
	var check func(cl, ro, n int) bool
	check = func(cl, ro, n int) bool {
		a := pre[cl+1][ro+1]
		b := pre[cl+1][max(0, ro+1-n)]
		c := pre[max(0, cl+1-n)][ro+1]
		d := pre[max(0, cl+1-n)][max(0, ro+1-n)]
		if a-b-c+d == n*n {
			return true
		}
		return false
	}
	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			for k := 1; k <= min(m, n); k++ {
				if check(i, j, k) {
					ans++
				}
			}
		}
	}
	return ans
}
