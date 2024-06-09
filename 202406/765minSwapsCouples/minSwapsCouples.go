package main

import (
	"fmt"

	. "outback/algorithm/common/unionfind"
)

func main() {
	fmt.Println(minSwapsCouples([]int{0, 2, 1, 3}))
}

func minSwapsCouples(row []int) int {
	n := len(row) / 2
	uf := NewRankUnionFind(n)

	for i := 0; i < n*2; i = i + 2 {
		uf.Union(row[i]/2, row[i+1]/2)
	}
	cnt := 0
	for i := 0; i < n; i++ {
		if uf.Find(i) == i {
			cnt++
		}
	}
	return n - cnt
}
