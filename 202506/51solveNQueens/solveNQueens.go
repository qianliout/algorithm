package main

import (
	"fmt"
)

func main() {
	ans := solveNQueens(4)
	for _, ch := range ans {
		fmt.Println(ch)
	}
}

func solveNQueens(n int) [][]string {

	ans := make([][]string, 0)
	path := make([][]byte, n)
	for i := range path {
		path[i] = make([]byte, n)
		for j := range path[i] {
			path[i][j] = '.'
		}
	}
	var dfs func(i int)
	dfs = func(i int) {
		if i >= n {
			ans = append(ans, gen(path))
			return
		}
		for k := 0; k < n; k++ {
			if check(path, i, k) {
				path[i][k] = 'Q'
				dfs(i + 1)
				path[i][k] = '.'
			}
		}
	}
	dfs(0)
	return ans
}

func gen(path [][]byte) []string {
	ans := make([]string, len(path))
	for i, ch := range path {
		ans[i] = string(ch)
	}
	return ans
}

func check(path [][]byte, i, j int) bool {
	for m := 0; m < i; m++ {
		if path[m][j] == 'Q' {
			return false
		}
	}
	for m := 0; m < j; m++ {
		if path[i][m] == 'Q' {
			return false
		}
	}

	for m, n := i, j; m >= 0 && n >= 0; m, n = m-1, n-1 {
		if path[m][n] == 'Q' {
			return false
		}
	}
	for m, n := i, j; m >= 0 && n < len(path); m, n = m-1, n+1 {
		if path[m][n] == 'Q' {
			return false
		}
	}

	return true
}
