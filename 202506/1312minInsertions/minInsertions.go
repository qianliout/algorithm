package main

import (
	"fmt"
)

func main() {
	fmt.Println(minInsertions("mbadm"))
	fmt.Println(minInsertions("zjveiiwvc"))
}

func minInsertions1(s string) int {
	n := len(s)
	mem := make([][]int, n+1)
	inf := 666

	for i := range mem {
		mem[i] = make([]int, n+1)
		for j := range mem[i] {
			mem[i][j] = inf
		}
	}
	// i->j 的结果值
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i == j {
			return 0
		}
		if i > j {
			return 0
			// zjveiiwvc
			// return inf // 这样返回就出错，为啥呢
		}
		if i >= n || j < 0 {
			return 0
		}
		if mem[i][j] != inf {
			return mem[i][j]
		}
		if s[i] == s[j] {
			ans := dfs(i+1, j-1)
			mem[i][j] = ans
			return ans
		}
		ans := min(dfs(i+1, j), dfs(i, j-1)) + 1
		mem[i][j] = ans
		return ans
	}

	return dfs(0, n-1)
}

func minInsertions2(s string) int {
	n := len(s)
	inf := 666
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, n+1)
		for j := range f[i] {
			f[i][j] = inf
		}
	}

	// f[i][j] = f[i+1][j-1]
	// f[i][j] = f[i+1[j] f[i][j-1]

	// 	i-1,j+1
	// f[i-1][j+1] = f[i][j]
	// f[i-1][j+1] = f[i[j+1] f[i-1][j]

	for i := n; i > 0; i-- {
		for j := 0; j < n; j++ {
			if i >= j {
				f[i][j] = 0
			}
			if s[i] == s[j] {
				f[i-1][j+1] = f[i][j]
			} else {
				f[i-1][j+1] = min(f[i][j+1], f[i-1][j]) + 1
			}
		}
	}
	return f[1][n]
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
