package main

import (
	"fmt"
)

func main() {
	edges := [][]int{{0, 1}, {0, 2}, {0, 3}, {0, 4}, {1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}}
	fmt.Println(getAncestors(5, edges)) //  [[],[0],[0,1],[0,1,2],[0,1,2,3]]
}

func getAncestors(n int, edges [][]int) [][]int {
	g := make([][]int, n)
	for _, ch := range edges {
		x, y := ch[0], ch[1]
		g[x] = append(g[x], y)
	}
	ans := make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, 0)
	}
	for i := 0; i < n; i++ {
		visit := make([]bool, n)
		res := dfs(g, i, visit)
		for _, ch := range res {
			ans[ch] = append(ans[ch], i)
		}
	}
	return ans
}

func dfs(g [][]int, start int, visit []bool) []int {
	nex := g[start]
	visit[start] = true
	ans := make([]int, 0)
	for _, ch := range nex {
		if visit[ch] {
			continue
		}
		ans = append(ans, ch)
		visit[ch] = true
		ans = append(ans, dfs(g, ch, visit)...)
	}
	return ans
}
