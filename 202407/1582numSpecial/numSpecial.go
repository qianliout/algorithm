package main

import (
	"fmt"
)

func main() {
	fmt.Println(numSpecial([][]int{{0, 0}, {0, 0}, {1, 0}}))
}

func numSpecial(mat [][]int) int {
	m, n := len(mat), len(mat[0])
	col := make([]int, m)
	row := make([]int, n)
	for i := 0; i < m; i++ {
		sum := 0
		for j := 0; j < n; j++ {
			sum += mat[i][j]
		}
		col[i] = sum
	}
	for j := 0; j < n; j++ {
		sum := 0
		for i := 0; i < m; i++ {
			sum += mat[i][j]
		}
		row[j] = sum
	}
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 1 && col[i] == 1 && row[j] == 1 {
				ans++
			}
		}
	}
	return ans
}
