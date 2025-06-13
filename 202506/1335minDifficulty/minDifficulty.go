package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(minDifficulty([]int{3}, 1))
}

func minDifficulty(job []int, d int) int {
	// 第i天，完成前j项工作的难度
	var dfs func(i, j int) int
	inf := 1 << 32
	n := len(job)
	mem := make([][]int, d+1)
	for i := range mem {
		mem[i] = make([]int, n+1)
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}
	dfs = func(i, j int) int {
		if i <= 0 || j < 0 {
			return inf
		}
		if i == 1 {
			if j >= n || j < 0 {
				return inf
			}
			return slices.Max(job[:j+1])
		}
		if mem[i][j] != -1 {
			return mem[i][j]
		}
		res := inf
		for k := j; k >= 0; k-- {
			res = min(res, slices.Max(job[k:j+1])+dfs(i-1, k-1))
		}
		mem[i][j] = res
		return res
	}
	a := dfs(d, n-1)
	if a >= inf {
		return -1
	}
	return a
}
