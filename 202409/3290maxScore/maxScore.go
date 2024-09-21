package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxScore([]int{3, 2, 5, 6}, []int{2, -6, 4, -5, -3, 2, -7}))
}

func maxScore1(a []int, b []int) int64 {
	inf := math.MinInt64 / 10
	n1, n2 := len(a), len(b)
	var dfs func(int, int) int
	mem := make([][]int, n1)
	for i := range mem {
		mem[i] = make([]int, n2)
		for j := range mem[i] {
			mem[i][j] = inf
		}
	}
	dfs = func(i int, j int) int {
		if i < 0 {
			return 0
		}
		if j < 0 {
			return inf
		}
		if mem[i][j] != inf {
			return mem[i][j]
		}
		mem[i][j] = max(dfs(i, j-1), dfs(i-1, j-1)+a[i]*b[j])
		return mem[i][j]
	}
	ans := dfs(n1-1, n2-1)
	return int64(ans)
}
func maxScore(a []int, b []int) int64 {
	inf := math.MinInt64 / 10
	n1, n2 := len(a), len(b)
	f := make([][]int, n1+1)
	for i := range f {
		f[i] = make([]int, n2+1)
		for j := range f[i] {
			f[i][j] = inf
		}
	}
	// 初值
	for i := 0; i < n2+1; i++ {
		f[0][i] = 0
	}
	for i := 0; i < n1; i++ {
		for j := 0; j < n2; j++ {
			f[i+1][j+1] = max(f[i+1][j], f[i][j]+a[i]*b[j])
		}
	}

	return int64(f[n1][n2])
}
