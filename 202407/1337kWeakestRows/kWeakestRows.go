package main

import (
	"sort"
)

func main() {

}

func kWeakestRows(mat [][]int, k int) []int {
	ans := make([]Week, 0)
	for i, ch := range mat {
		ans = append(ans, Week{
			Cnt: sum(ch),
			Idx: i,
		})
	}
	sort.Slice(ans, func(i, j int) bool {
		if ans[i].Cnt < ans[j].Cnt {
			return true
		} else if ans[i].Cnt > ans[j].Cnt {
			return false
		}
		return ans[i].Idx < ans[j].Idx
	})
	res := make([]int, 0)
	for i, ch := range ans {
		if i > k {
			break
		}
		res = append(res, ch.Idx)
	}
	return res
}

func sum(arr []int) int {
	ans := 0
	for _, ch := range arr {
		ans += ch
	}
	return ans
}

type Week struct {
	Cnt int
	Idx int
}
