package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(kthLargestValue([][]int{{5, 2}, {1, 6}}, 1))
}

func kthLargestValue(matrix [][]int, k int) int {
	n, m := len(matrix), len(matrix[0])
	if k > n*m {
		return 0
	}
	pre := make([][]int, n+1)
	for i := range pre {
		pre[i] = make([]int, m+1)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			pre[i+1][j+1] = pre[i+1][j] ^ pre[i][j+1] ^ pre[i][j] ^ matrix[i][j]
		}
	}
	ans := make([]int, m*n)
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			ans = append(ans, pre[i][j])
		}
	}
	sort.Slice(ans, func(i, j int) bool { return ans[j] < ans[i] })

	return ans[k-1]
}
