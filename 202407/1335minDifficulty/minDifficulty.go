package main

import (
	"fmt"
	"math"
	"slices"
)

func main() {
	fmt.Println(minDifficulty([]int{11, 111, 22, 222, 33, 333, 44, 444}, 6))
	fmt.Println(minDifficulty([]int{6, 5, 4, 3, 2, 1}, 2))
	fmt.Println(minDifficulty([]int{1, 1, 1}, 3))

}

func minDifficulty(job []int, d int) int {
	n := len(job)
	// 第i天(从一开始) 完成 j 项工作的结果
	inf := math.MaxInt / 2
	var dfs func(i, j int) int
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
		// 这一天能完成的数
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
