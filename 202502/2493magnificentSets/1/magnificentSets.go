package main

import (
	"fmt"
)

func main() {
	fmt.Println(magnificentSets(6, [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 2}, {1, 4}}))
	fmt.Println(magnificentSets(6, [][]int{{1, 2}, {1, 4}, {1, 5}, {2, 6}, {2, 3}, {4, 6}})) // 4
	fmt.Println(magnificentSets(3, [][]int{{1, 2}, {2, 3}, {3, 1}}))
}

func magnificentSets(n int, edges [][]int) int {
	g := make([][]int, n+1)
	for _, ch := range edges {
		x, y := ch[0]-1, ch[1]-1
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	var bipartite func(start, c int) bool
	color := make([]int, n)
	ti := make([]int, n)

	clo := 1
	bfs := func(start int) int {
		dep := 0
		ti[start] = clo
		q := []int{start}
		for len(q) > 0 {
			dep++
			lev := make([]int, 0)
			for _, x := range q {
				for _, y := range g[x] {
					if ti[y] != clo {
						lev = append(lev, y)
						ti[y] = clo
					}
				}
			}
			q = lev
		}
		return dep
	}

	nodes := make([]int, 0)
	bipartite = func(start, c int) bool {
		nodes = append(nodes, start)
		color[start] = c
		for _, y := range g[start] {
			if color[y] == c || (color[y] == 0 && !bipartite(y, -c)) {
				return false
			}
		}
		return true
	}

	ans := 0
	for i, c := range color {
		if c != 0 {
			// 说明在其他步中已经访问过
			continue
		}
		nodes = make([]int, 0)
		if !bipartite(i, 1) {
			return -1
		}
		th := 0
		for _, x := range nodes {
			// 每次调用时的时钟不一样
			clo++
			th = max(th, bfs(x))
		}
		ans += th
	}
	return ans
}
