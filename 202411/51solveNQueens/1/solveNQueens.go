package main

import (
	"fmt"
	"strings"
)

func main() {
	// fmt.Println(solveNQueens(4))
	fmt.Println(solveNQueens(1))
	fmt.Println(solveNQueens(8))
}

// 写法是错的
func solveNQueens(n int) [][]string {
	col := make([]int, n)
	ans := make([][]string, 0)
	// 只检查对角线，因下 dfs 中会检查行和列
	check := func(r, c int) bool {

		for r0 := 0; r0 < r; r0++ {
			c0 := col[r0]
			if r0+c0 == r+c || r0-c0 == r-c {
				return false
			}
		}
		return true
	}
	gen := func(s []int) []string {
		a := make([]string, n)
		for i, j := range s {
			b := strings.Repeat(".", j) + "Q" + strings.Repeat(".", max(0, n-j-1))
			a[i] = b
		}
		return a
	}

	// 其中 s 表示棋盘的中竖列还有那些可用
	var dfs func(c int, s []int)
	dfs = func(i int, s []int) {
		if i >= n {
			ans = append(ans, gen(s)) //
			return
		}
		for j := 0; j < n; j++ {
			if s[j] > 0 {
				continue
			}
			if !check(i, j) {
				continue
			}
			col[i] = j
			s[i] = 1
			dfs(i+1, s)
			// s[i] = 0
			// col[i] = 0
		}
	}
	s := make([]int, n)
	dfs(0, s)
	return ans
}
