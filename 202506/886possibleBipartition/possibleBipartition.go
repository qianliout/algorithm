package main

import (
	. "outback/algorithm/common/unionfind"
)

func main() {

}

func possibleBipartition(n int, dislikes [][]int) bool {
	uf := NewSizeUnionFind(n)
	g := make([][]int, n)
	for _, ch := range dislikes {
		x, y := ch[0]-1, ch[1]-1
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	for x, ch := range g {
		for _, y := range ch {
			if uf.IsConnected(x, y) {
				return false
			}
			uf.Union(y, ch[0])
		}
	}

	return true
}
