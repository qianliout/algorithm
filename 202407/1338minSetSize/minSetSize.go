package main

import (
	"sort"
)

func main() {

}

func minSetSize(arr []int) int {
	cnt := make(map[int]int)
	for _, ch := range arr {
		cnt[ch]++
	}
	pairs := make([]pair, 0)
	for k, v := range cnt {
		pairs = append(pairs, pair{Num: k, Cnt: v})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Cnt > pairs[j].Cnt
	})
	n := len(arr)
	sub := n
	ans := 0
	for _, pa := range pairs {
		if sub <= n/2 {
			break
		}
		ans++
		sub -= pa.Cnt
	}

	return ans
}

type pair struct {
	Num int
	Cnt int
}
