package main

import (
	"fmt"
)

func main() {
	fmt.Println(minCostConnectPoints([][]int{{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0}})) // 20
}

func minCostConnectPoints(points [][]int) int {
	n := len(points)
	visit := make([]bool, n) // 标记点是否已加入MST(Minimum Spanning Tree,最小生成树)
	dist := make([]int, n)   // 每个点到MST的最短距离,表示点j到整个MST的最短距离，不是到某个特定点的距离
	inf := 1 << 31

	for i := range dist {
		dist[i] = inf
	}
	dist[n-1] = 0 // 随机一个节点
	for i := 0; i < n; i++ {
		u := -1
		for j := 0; j < n; j++ {
			if !visit[j] && (u == -1 || dist[j] < dist[u]) {
				u = j
			}
		}
		if u == -1 {
			break
		}
		/*
			在Prim算法中，dist[j] 表示点j到整个MST的最短距离，不是到某个特定点的距离。
				举个例子：
				假设我们有4个点：A、B、C、D
				初始：MST = {A}
				第一轮：选择离A最近的点B加入MST，MST = {A, B}
				关键：现在需要重新计算C和D到MST的距离
				对于点C：
				之前 dist[C] = C到A的距离
				现在需要比较：C到A的距离 vs C到B的距离
				dist[C] = min(原来的dist[C], C到B的距离)
		*/
		visit[u] = true
		for j := 0; j < n; j++ {
			if !visit[j] {
				dist[j] = min(dist[j], abs(points[j][0]-points[u][0])+abs(points[j][1]-points[u][1]))
			}
		}
	}
	ans := 0
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
