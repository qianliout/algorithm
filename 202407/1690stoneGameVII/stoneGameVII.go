package main

import (
	"fmt"
)

func main() {
	fmt.Println(stoneGameVII([]int{5, 3, 1, 4, 2}))
}

func stoneGameVII(stones []int) int {
	n := len(stones)
	sum := make([]int, n+1)
	for i, ch := range stones {
		sum[i+1] = sum[i] + ch
	}
	mem := make([][]int, n)
	for i := range mem {
		mem[i] = make([]int, n)
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}

	var dfs func(i, j int) int

	dfs = func(i, j int) int {
		// 边界
		// if i > j || i < 0 || j < 0 || j >= n || i >= n {
		// 	return math.MinInt / 10
		// }
		if i == j {
			return 0
		}
		if mem[i][j] != -1 {
			return mem[i][j]
		}
		res := 0
		res = max(res, sum[j+1]-sum[i+1]-dfs(i+1, j))
		res = max(res, sum[j]-sum[i]-dfs(i, j-1))
		mem[i][j] = res
		return res
	}
	res := dfs(0, n-1)
	return res
}
