package main

import (
	"fmt"
)

func main() {
	fmt.Println(minDistance("horse", "ros"))
	fmt.Println(minDistance2("horse", "ros"))
}

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	f := make([][]int, m+5)
	for i := range f {
		f[i] = make([]int, n+5)
	}
	for i := 0; i <= m; i++ {
		f[i][0] = i
	}
	for j := 0; j <= n; j++ {
		f[0][j] = j
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				f[i][j] = f[i-1][j-1]
			} else {
				f[i][j] = min(f[i-1][j], f[i][j-1], f[i-1][j-1]) + 1
			}
		}
	}
	return f[m][n]
}

func minDistance2(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	var dfs func(i, j int) int
	inf := m*n + 10
	dfs = func(i, j int) int {
		if i < 0 || j < 0 {
			return inf
		}
		if i > m || j > n {
			return inf
		}
		if i == 0 {
			return j
		}
		if j == 0 {
			return i
		}
		if word1[i-1] == word2[j-1] {
			return dfs(i-1, j-1)
		}
		ans := min(dfs(i-1, j), dfs(i, j-1), dfs(i-1, j-1)) + 1
		return ans
	}
	ans := dfs(m, n)
	return ans
}
