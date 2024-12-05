package main

import (
	"math"
)

func main() {

}

func minimumDeleteSum(s1 string, s2 string) int {

	m, n := len(s1), len(s2)
	pre1, pre2 := make([]int, m+1), make([]int, n+1)
	for i, ch := range s1 {
		pre1[i+1] = pre1[i] + int(ch)
	}
	for j, ch := range s2 {
		pre2[j+1] = pre2[j] + int(ch)
	}

	inf := math.MaxInt / 10
	f := make([][]int, m+1)
	for i := range f {
		f[i] = make([]int, n+1)
		for j := range f[i] {
			f[i][j] = inf
		}
	}
	// 初值
	for i := 0; i <= m; i++ {
		f[i][0] = pre1[i]
	}
	for j := 0; j <= n; j++ {
		f[0][j] = pre2[j]
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s1[i-1] == s2[j-1] {
				f[i][j] = f[i-1][j-1]
			} else {
				f[i][j] = min(f[i-1][j]+int(s1[i-1]), f[i][j-1]+int(s2[j-1]))
			}
		}
	}
	return f[m][n]
}

func minimumDeleteSum1(s1 string, s2 string) int {
	m, n := len(s1), len(s2)
	var dfs func(i, j int) int
	pre1, pre2 := make([]int, m+1), make([]int, n+1)
	for i, ch := range s1 {
		pre1[i+1] = pre1[i] + int(ch)
	}
	for j, ch := range s2 {
		pre2[j+1] = pre2[j] + int(ch)
	}

	inf := math.MaxInt / 10
	dfs = func(i, j int) int {
		if i < 0 || j < 0 || i > m || j > n {
			return inf
		}

		if i == 0 {
			return pre2[j]
		}
		if j == 0 {
			return pre1[i]
		}
		if s1[i-1] == s2[j-1] {
			return dfs(i-1, j-1)
		}
		return min(dfs(i-1, j)+int(s1[i-1]), dfs(i, j-1)+int(s2[j-1]))
	}
	ans := dfs(m, n)
	return ans
}
