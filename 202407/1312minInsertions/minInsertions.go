package main

import (
	"fmt"
)

func main() {
	fmt.Println(minInsertions("zzazz"))
}

func minInsertions1(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1]
			} else {
				dp[i][j] = min(dp[i+1][j], dp[i][j-1]) + 1
			}
		}
	}
	return dp[0][n-1]
}

func minInsertions2(s string) int {
	n := len(s)
	mem := make([][]int, n)
	for i := range mem {
		mem[i] = make([]int, n)
	}
	// dfs 的作用是s[i:j](包括 j) 里最长文件子序列
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i > j {
			return 0
		}
		if i == j {
			return 1
		}
		if mem[i][j] > 0 {
			return mem[i][j]
		}
		res := 0
		if s[i] == s[j] {
			res = dfs(i+1, j-1) + 2
		} else {
			res = max(dfs(i+1, j), dfs(i, j-1))
		}
		mem[i][j] = res

		return res
	}
	// 字符串长度是 n，最长回文件子序列是 A 那么,需要插入的字符就是 n-A
	return n - dfs(0, n-1)
}
func minInsertions(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		dp[i][i] = 1
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return n - dp[0][n-1]
}
