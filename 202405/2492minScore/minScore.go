package main

import (
	"math"
)

func main() {

}

type edge struct{ to, d int }

// 城市 1 和城市 n 之间的所有路径的 最小 分数。
// 测试数据保证城市 1 和城市n 之间 至少 有一条路径
// 其实就是求从1开始的路径中，分数最小的两个城市的分数

func minScore(n int, roads [][]int) int {
	graph := make([][]edge, n)
	for _, ch := range roads {
		x, y, d := ch[0]-1, ch[1]-1, ch[2]
		graph[x] = append(graph[x], edge{y, d})
		graph[y] = append(graph[y], edge{x, d})
	}
	ans := math.MaxInt32
	visit := make([]bool, n)

	dfs(graph, 0, visit, &ans)
	return ans
}

func dfs(graph [][]edge, start int, visit []bool, ans *int) {
	visit[start] = true
	nex := graph[start]
	for _, ch := range nex {
		*ans = min(*ans, ch.d)
		if visit[ch.to] {
			continue
		}
		dfs(graph, ch.to, visit, ans)
	}
}
