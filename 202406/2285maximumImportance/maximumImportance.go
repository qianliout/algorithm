package main

import (
	"sort"
)

func main() {

}

func maximumImportance(n int, roads [][]int) int64 {
	g := make([]pair, n)
	for i := 0; i < n; i++ {
		g[i] = pair{node: i, path: make([]int, 0)}
	}
	for _, ch := range roads {
		x, y := ch[0], ch[1]
		g[x].path = append(g[x].path, y)
		g[y].path = append(g[y].path, x)
	}

	sort.Slice(g, func(i, j int) bool { return len(g[i].path) > len(g[j].path) })
	value := make(map[int]int)
	start := n
	for _, p := range g {
		p.value = start
		value[p.node] = start
		start--
	}
	ans := 0
	for _, p := range g {
		for _, pa := range p.path {
			ans += value[pa] + p.value
		}
	}
	return int64(ans)
}

type pair struct {
	node  int
	value int
	path  []int
}
