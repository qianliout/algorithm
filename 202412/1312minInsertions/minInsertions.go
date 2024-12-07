package main

import (
	"fmt"
)

func main() {
	fmt.Println(minInsertions("zjveiiwvc"))
}

func minInsertions1(s string) int {
	n := len(s)
	var dfs func(i, j int) int
	// inf := math.MaxInt >> 2
	mem := make([][]int, n)
	for i := range mem {
		mem[i] = make([]int, n)
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}

	dfs = func(i, j int) int {
		if i > j || j < 0 || i < 0 {
			return 0
		}

		if i == j {
			return 0
		}
		if mem[i][j] != -1 {
			return mem[i][j]
		}
		if s[i] == s[j] {
			return dfs(i+1, j-1)
		}

		ans := min(dfs(i+1, j), dfs(i, j-1)) + 1
		mem[i][j] = ans
		return ans
	}

	return dfs(0, n-1)
}
func minInsertions(s string) int {
	n := len(s)
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}

	for i := n; i > 0; i-- {
		for j := i + 1; j <= n; j++ {
			if s[i-1] == s[j-1] {
				f[i][j] = f[i+1][j-1]
			} else {
				f[i][j] = min(f[i+1][j], f[i][j-1]) + 1
			}

		}
	}
	return f[1][n]
}
