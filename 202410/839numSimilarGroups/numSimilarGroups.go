package main

import (
	"fmt"
	. "outback/algorithm/common/unionfind"
)

func main() {
	fmt.Println(numSimilarGroups([]string{"tars", "rats", "arts", "star"}))
}

func numSimilarGroups(strs []string) int {
	n := len(strs)
	uf := NewRankUnionFind(n)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			// 两个写法都是对的
			// if !uf.IsConnected(i, j) && similar(strs[i], strs[j]) {
			if similar(strs[i], strs[j]) {
				uf.Union(i, j)
			}
		}
	}
	return uf.Count
}
func similar(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	cnt := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			cnt++
		}
	}
	return cnt == 2 || cnt == 0
}
