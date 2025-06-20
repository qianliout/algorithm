package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minCostConnectPoints([][]int{{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0}})) // 20
}

func minCostConnectPoints(points [][]int) int {
	inf := math.MaxInt / 10
	n := len(points)
	ans := 0
	visit := make([]bool, n) // 标记点是否已加入MST(Minimum Spanning Tree,最小生成树)
	dist := make([]int, n)   // 每个点到MST的最短距离
	for i := range dist {
		dist[i] = inf
	}
	dist[2] = 0
	for i := 0; i < n; i++ {
		u := -1
		// 找到未访问点中距离MST最近的点
		for j := 0; j < n; j++ {
			if !visit[j] && (u == -1 || dist[u] > dist[j]) {
				u = j
				// 因为需要遍历所有未访问的点来找到距离最小的那个。所以不能 break
				// 这里不能写break 这一组结果可以测试： [[0,0],[1,1],[1,0],[-1,1]]
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

/*
给你一个points 数组，表示 2D 平面上的一些点，其中 points[i] = [xi, yi] 。
连接点 [xi, yi] 和点 [xj, yj] 的费用为它们之间的 曼哈顿距离 ：|xi - xj| + |yi - yj| ，其中 |val| 表示 val 的绝对值。
请你返回将所有点连接的最小总费用。只有任意两点之间 有且仅有 一条简单路径时，才认为所有点都已连接。
*/
