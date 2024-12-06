package main

import (
	"math"
)

func main() {

}

func numDistinct(s string, t string) int {
	m, n := len(s), len(t)
	mod := int(math.Pow(10, 9) + 7)
	f := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		f[i] = make([]int, n+1)
	}
	// 初值
	for i := 0; i <= m; i++ {
		f[i][0] = 1
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s[i-1] == t[j-1] {
				f[i][j] = f[i-1][j-1] + f[i-1][j]
				f[i][j] = f[i][j] % mod
			} else {
				f[i][j] = f[i-1][j]
			}
		}
	}
	return f[m][n] % mod
}

func numDistinct1(s string, t string) int {
	m, n := len(s), len(t)
	mod := int(math.Pow(10, 9) + 7)
	var dfs func(i, j int) int
	mem := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		mem[i] = make([]int, n+1)
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}

	dfs = func(i, j int) int {
		if j <= 0 {
			return 1
		}
		if i <= 0 {
			return 0
		}

		if mem[i][j] != -1 {
			return mem[i][j]
		}

		if s[i-1] == t[j-1] {
			ans := dfs(i-1, j-1) + dfs(i-1, j)
			mem[i][j] = ans % mod
			return ans % mod
		}
		ans := dfs(i-1, j)
		mem[i][j] = ans % mod
		return ans % mod
	}
	ans := dfs(m, n)
	return ans % mod
}
