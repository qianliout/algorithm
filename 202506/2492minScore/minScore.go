package main

import (
	"math"
)

func main() {

}

type edge struct{ to, d int }

func minScore(n int, roads [][]int) int {
	graph := make([][]edge, n)
	for _, ch := range roads {
		x, y, d := ch[0]-1, ch[1]-1, ch[2]
		graph[x] = append(graph[x], edge{y, d})
		graph[y] = append(graph[y], edge{x, d})
	}
	ans := math.MaxInt32
	visit := make([]bool, n)
	var dfs func(start int)
	dfs = func(start int) {
		visit[start] = true
		nex := graph[start]
		for _, ch := range nex {
			ans = min(ans, ch.d)
			if visit[ch.to] {
				continue
			}
			dfs(ch.to)
		}
	}
	dfs(0)
	return ans
}

// func dfs(graph [][]edge, start int, visit []bool, ans *int) {
// 	visit[start] = true
// 	nex := graph[start]
// 	for _, ch := range nex {
// 		*ans = min(*ans, ch.d)
// 		if visit[ch.to] {
// 			continue
// 		}
// 		dfs(graph, ch.to, visit, ans)
// 	}
// }
