package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findShortestCycle(6, [][]int{{4, 2}, {5, 1}, {5, 0}, {0, 3}, {5, 2}, {1, 4}, {1, 3}, {3, 4}}))
}

func findShortestCycle(n int, edges [][]int) int {
	g := make([][]int, n)
	for i := range edges {
		x, y := edges[i][0], edges[i][1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	ans := math.MaxInt32
	for i := 0; i < n; i++ {
		b := bfs(n, g, i)
		ans = min(b, ans)
	}
	if ans == math.MaxInt32 {
		return -1
	}
	return ans
}

func bfs(n int, g [][]int, start int) int {

	dis := make([]int, n)
	for i := range dis {
		dis[i] = -1
	}
	dis[start] = 0
	type pair struct{ idx, pa int } // pa 表示 idx 是从那一个节点走过来的
	queue := make([]pair, 0)
	queue = append(queue, pair{idx: start, pa: -1})
	ans := math.MaxInt32
	for len(queue) > 0 {
		no := queue[0]
		x, pa := no.idx, no.pa
		queue = queue[1:]
		for _, y := range g[x] {
			if dis[y] < 0 {
				dis[y] = dis[x] + 1
				queue = append(queue, pair{y, x})
			} else if y != pa { // 说明以前有节点返问到这个节点
				ans = min(ans, dis[x]+dis[y]+1)
			}
		}
	}
	return ans
}
