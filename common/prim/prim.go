package prim

import (
	"math"
)

func minCostConnectPoints(points [][]int) int {
	inf := math.MaxInt / 10
	n := len(points)
	ans := 0
	visit := make([]bool, n)
	dist := make([]int, n)
	for i := range dist {
		dist[i] = inf
	}
	dist[0] = 0 // 随便一个点入队
	for i := 0; i < n; i++ {
		u := -1
		for j := 0; j < n; j++ {
			if !visit[j] && (u == -1 || dist[u] > dist[j]) {
				u = j
				// 这里不能写break 这一组结果可以测试： [[0,0],[1,1],[1,0],[-1,1]]
				// 这里的目的是找一条最小的边，所以当找到第一个 j 的时候就不能跳出循环，需要继续找
				// break
			}
		}
		if u == -1 {
			break
		}
		visit[u] = true
		for j := 0; j < n; j++ {
			if !visit[j] {
				nd := abs(points[j][0]-points[u][0]) + abs(points[j][1]-points[u][1])
				dist[j] = min(dist[j], nd)
			}
		}
	}

	for _, ch := range dist {
		ans += ch
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
