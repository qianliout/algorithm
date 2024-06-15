package main

import (
	"fmt"
)

func main() {
	fmt.Println(matrixBlockSum([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 1))
}

/*
给你一个 m x n 的矩阵 mat 和一个整数 k ，请你返回一个矩阵 answer ，其中每个 answer[i][j] 是所有满足下述条件的元素 mat[r][c] 的和：

    i - k <= r <= i + k,
    j - k <= c <= j + k 且
    (r, c) 在矩阵内。


*/

// 数据不多，直接模拟
func matrixBlockSum(mat [][]int, k int) [][]int {
	n, m := len(mat), len(mat[0])
	pre := make([][]int, n+1)
	for i := range pre {
		pre[i] = make([]int, m+1)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			pre[i+1][j+1] = pre[i][j+1] + pre[i+1][j] - pre[i][j] + mat[i][j]
		}
	}

	ans := make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			a := pre[min(n, i+k+1)][min(m, j+k+1)]
			b := pre[min(n, i+k+1)][max(0, j-k)]
			c := pre[max(0, i-k)][min(m, j+k+1)]
			d := pre[max(0, i-k)][max(0, j-k)]
			ans[i][j] = a - b - c + d
		}
	}
	return ans
}
