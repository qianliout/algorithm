package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(connectTwoGroups([][]int{{1, 3, 5}, {4, 1, 1}, {1, 5, 3}}))
}

func connectTwoGroups(cost [][]int) int {
	n := len(cost)
	m := len(cost[0])
	minCost := make([]int, m)
	inf := math.MaxInt / 10
	for j := 0; j < m; j++ {
		minCost[j] = inf
		for i := 0; i < n; i++ {
			minCost[j] = min(minCost[j], cost[i][j])
		}
	}

	mem := make([][]int, n)
	for i := range mem {
		mem[i] = make([]int, 1<<m)
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i < 0 {
			// B中没有和 A 连接的，就去找一个最小的点连上
			ans := 0
			for k := 0; k < m; k++ {
				if j>>k&1 == 1 {
					ans += minCost[k]
				}
			}
			return ans
		}
		if mem[i][j] != -1 {
			return mem[i][j]
		}
		res := inf
		for k, c := range cost[i] {
			res = min(res, dfs(i-1, j&^(1<<k))+c)
		}
		mem[i][j] = res
		return res
	}
	res := dfs(n-1, 1<<m-1)
	return res
}
