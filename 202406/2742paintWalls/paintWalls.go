package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(paintWalls([]int{1, 2, 3, 2}, []int{1, 2, 3, 2}))
}

func paintWalls(cost []int, time []int) int {
	inf := math.MaxInt / 2
	var dfs func(i, j int) int

	// i表示i个墙
	// j 表示，付费时间-已使用的免费时间
	// 这里要注意的是 j 有可能是负数，因为可以先做免费的，后面再做付费的
	n := len(cost)
	mem := make([]map[int]int, n)
	for i := range mem {
		mem[i] = make(map[int]int)
	}

	dfs = func(i, j int) int {
		if j >= i+1 {
			// 下标是 i 那说明还有 i+1面墙，因为下标是从0开始的
			// 剩余的都可以免费
			return 0
		}
		if i < 0 {
			// 刷完了
			return inf
		}
		if va, ok := mem[i][j]; ok {
			return va
		}
		res := inf
		res = min(res, dfs(i-1, j+time[i])+cost[i])
		res = min(res, dfs(i-1, j-1))
		mem[i][j] = res
		return res
	}

	return dfs(n-1, 0)
}
