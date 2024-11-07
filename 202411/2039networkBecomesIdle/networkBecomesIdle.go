package main

import (
	"fmt"
)

func main() {
	// edges = [[0,1],[1,2]], patience = [0,2,1]
	fmt.Println(networkBecomesIdle([][]int{{0, 1}, {1, 2}}, []int{0, 2, 1}))
}

func networkBecomesIdle(edges [][]int, patience []int) int {
	n := len(patience)
	g := make([][]int, n)
	for _, ch := range edges {
		x, y := ch[0], ch[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	// bfs求出各个点到0号点的最短路径
	dist := make([]int, n)
	for i := range dist {
		dist[i] = n + 1
	}
	dist[0] = 0
	queue := []pair{{0, 0}}
	for len(queue) > 0 {
		lev := make([]pair, 0)
		for _, no := range queue {
			for _, y := range g[no.idx] {
				// 这里一定要判断，因为是 bfs，所以上一步一定是最短路
				if dist[y] == n+1 {
					dist[y] = min(dist[y], no.dis+1)
					lev = append(lev, pair{y, dist[y]})
				}
			}
		}
		queue = lev
	}
	ans := 0
	for i := 1; i < n; i++ {
		d := dist[i] * 2
		p := patience[i]
		if p > 0 {
			ans = max(ans, (d-1)/p*p+d)
		}
	}

	return ans + 1
}

type pair struct {
	idx, dis int
}
