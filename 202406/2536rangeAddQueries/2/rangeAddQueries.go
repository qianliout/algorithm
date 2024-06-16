package main

import (
	"fmt"
)

func main() {
	// fmt.Println(rangeAddQueries(3, [][]int{{1, 1, 2, 2}, {0, 0, 1, 1}}))
	fmt.Println(rangeAddQueries(2, [][]int{{0, 0, 1, 1}}))
}

func rangeAddQueries(n int, queries [][]int) [][]int {
	ans := make([]int, n*n+2)
	for _, ch := range queries {
		pairs := gen(ch[0], ch[1], ch[2], ch[3], n)
		for _, p := range pairs {
			ans[p.start+1]++
			ans[p.end+2]--
		}
	}
	for i := 1; i < len(ans); i++ {
		ans[i] += ans[i-1]
	}
	// 复原
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	for i := 1; i <= n*n; i++ {
		a := i - 1
		matrix[a/n][a%n] = ans[a]
	}
	return matrix
}

// 二维变一维
func exp(a, b, rowC int) int {
	return a*rowC + b
}

type pair struct {
	start, end int
}

func gen(col1, row1, col2, row2, rowC int) []pair {
	ans := make([]pair, 0)
	// leetcode中的 row 和col 和我的理解是反的
	// query[i] = [row1i, col1i, row2i, col2i]
	for i := col1; i <= col2; i++ {
		start := exp(i, row1, rowC)
		end := exp(i, row2, rowC)
		ans = append(ans, pair{start, end})
	}

	return ans
}
