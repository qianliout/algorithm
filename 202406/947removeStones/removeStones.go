package main

import (
	. "outback/algorithm/common/unionfind"
)

func main() {

}

func removeStones(stones [][]int) int {
	n := 20005
	uf := NewRankUnionFind(n)
	for _, ch := range stones {
		x, y := ch[0], ch[1]+10003
		uf.Union(x, y)
	}

	exist := make(map[int]bool)
	for _, ch := range stones {
		x, y := ch[0], ch[1]+10003
		exist[uf.Find(x)] = true
		exist[uf.Find(y)] = true
	}
	return len(stones) - len(exist)
}
