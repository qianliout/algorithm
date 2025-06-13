package main

import (
	. "outback/algorithm/common/unionfind"
)

func main() {

}

func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
	uf := NewSizeUnionFind(n + 10)
	ans := make([]int, 0)
	for _, ch := range edges {
		x, y := ch[0], ch[1]
		if uf.IsConnected(x, y) {
			ans = []int{x, y}
		}
		uf.Union(x, y)
	}
	return ans
}
