package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(diagonalSort([][]int{{37, 98, 82, 45, 42}}))
}

func diagonalSort(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	g := make([][]int, m+n)
	// 具体地，我们记矩阵的行数为 m，列数为 n。由于同一条对角线上的任意两个元素 (i1,j1) 和 (j2,j2) 满足 j1-i1==j2-i2
	// 我们可以根据 j−i 的值来确定每条对角线。为了保证值为正数，我们加上一个偏移量 m，即 m−i+j。
	for i := 0; i < m; i++ {
		for j, ch := range mat[i] {
			g[m-i+j] = append(g[m-i+j], ch)
		}
	}
	for i := 0; i < len(g); i++ {
		sort.Ints(g[i])
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// k := len(g[m-i+j])
			mat[i][j] = g[m-i+j][0]
			g[m-i+j] = g[m-i+j][1:]
		}
	}
	return mat
}
