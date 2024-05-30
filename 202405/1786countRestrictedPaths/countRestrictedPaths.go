package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(countRestrictedPaths(5, [][]int{{1, 2, 3}, {1, 3, 3}, {2, 3, 1}, {1, 4, 2}, {5, 2, 2}, {3, 5, 1}, {5, 4, 10}}))
}

func countRestrictedPaths(n int, edges [][]int) int {
	inf := math.MaxInt
	g := make([][]int, n+1)
	for i := range g {
		g[i] = make([]int, n+1)
		for j := range g[i] {
			g[i][j] = inf
		}
	}
	for _, ch := range edges {
		x, y, z := ch[0], ch[1], ch[2]
		g[x][y] = z
		g[y][x] = z
	}
	dis := make([]int, n+1)
	visit := make([]bool, n+1)
	dis[n] = 0
	dp := make([]int, n+1)
	dp[n] = 1
	for {
		x := -1
		for i, ok := range visit {
			if !ok && (x < 0 || dis[x] > dis[i]) {
				x = i
			}
		}
		if x < 0 {
			break
		}
		visit[x] = true
		// 不失一般性，当我们要求 dp[i] 的时候，其实找的所有满足「与点 i 相连，且最短路比点 i 要小的点 j」，
		// 符合条件的点 j 有很多个，将所有的 dp[j] 累加即是 dp[i]。
		// dp[i] = dp[i - 1] if dis[i-1] < dis[i]
		for y, d := range g[x] {
			if d >= inf {
				continue
			}
			dis[y] = min(dis[y], dis[x]+d)
		}
	}
	return 0
}
