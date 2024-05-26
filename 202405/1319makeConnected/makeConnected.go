package main

import (
	"fmt"

	. "outback/algorithm/common/unionfind"
)

func main() {
	fmt.Println(makeConnected(4, [][]int{{0, 1}, {0, 2}, {1, 2}}))
}

func makeConnected(n int, connections [][]int) int {
	uf := NewRankUnionFind(n)
	reset := 0
	for _, ch := range connections {
		if uf.IsConnected(ch[0], ch[1]) {
			reset++
		}
		uf.Union(ch[0], ch[1])
	}
	if uf.Count-1 > reset {
		return -1
	}
	return uf.Count - 1
}
