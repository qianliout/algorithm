package main

import (
	"math"
)

func main() {

}

func mostProfitablePath(edges [][]int, bob int, amount []int) int {
	n := len(edges) + 1
	g := make([][]int, n)

	for _, ch := range edges {
		x, y := ch[0], ch[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	inf := math.MaxInt / 10
	bobTime := make([]int, n)
	for i := 0; i < n; i++ {
		bobTime[i] = inf
	}
	var dfs1 func(i, fa, ti int) bool
	dfs1 = func(i, fa, ti int) bool {
		if i == 0 {
			bobTime[i] = ti
			return true
		}
		for _, nx := range g[i] {
			if nx != fa && dfs1(nx, i, ti+1) {
				bobTime[i] = ti
				return true
			}
		}
		return false
	}
	dfs1(bob, -1, 0)

	ans := math.MinInt32
	g[0] = append(g[0], -1)
	var dfs2 func(x, fa, aliceTime, total int)
	dfs2 = func(x, fa, aliceTime, total int) {
		if aliceTime < bobTime[x] {
			total += amount[x]
		} else if aliceTime == bobTime[x] {
			total += amount[x] / 2
		}
		if len(g[x]) == 1 {
			ans = max(ans, total)
			return
		}
		for _, nx := range g[x] {
			if nx != fa {
				dfs2(nx, x, aliceTime+1, total)
			}
		}
	}
	dfs2(0, -1, 0, 0)
	return ans
}
