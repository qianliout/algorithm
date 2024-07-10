package main

import (
	"fmt"
)

func main() {
	fmt.Println(minTime(7, [][]int{{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6}}, []bool{false, false, true, false, true, true, false}))
}

func minTime(n int, edges [][]int, hasApple []bool) int {
	g := make([][]int, n)
	for _, ch := range edges {
		x, y := ch[0], ch[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	// 标识有苹果的路
	var dfs1 func(i, fa int) bool
	dfs1 = func(i, fa int) bool {
		if i >= n {
			return false
		}
		if hasApple[i] {
			return true
		}

		res := false
		for _, nx := range edges[i] {
			if nx == fa {
				continue
			}
			res = res || dfs1(nx, i)
		}
		hasApple[i] = res
		return res
	}

	dfs1(0, -1)
	fmt.Println(hasApple)
	var dfs2 func(i, fa int) int

	mem := make([][]int, n+1)
	for i := range mem {
		mem[i] = make([]int, n+1)
	}
	// ans := 0
	dfs2 = func(i, fa int) int {
		if i >= n {
			return 0
		}
		if !hasApple[i] {
			return 0
		}
		res := 0
		for _, nx := range edges[i] {
			if nx == fa {
				continue
			}
			if !hasApple[nx] {
				continue
			}
			res += dfs2(nx, i) + 2
		}
		mem[i][fa] = res
		return res
	}
	res := dfs2(0, n)
	return res
}
