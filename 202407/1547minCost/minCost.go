package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(minCost(7, []int{1, 3, 4, 5}))
}

func minCost(n int, cuts []int) int {
	sort.Ints(cuts)
	var dfs func(le, ri int) int
	inf := math.MaxInt / 10
	// n 会很大，但是 cuts 不大，所以要用 map 做离散化
	mem := make(map[int]map[int]int)
	dfs = func(le, ri int) int {
		if le >= ri {
			return 0
		}
		if mem[le] != nil {
			if va, ok := mem[le][ri]; ok {
				return va
			}
		}
		ret := inf
		for _, ch := range cuts {
			if ch > le && ch < ri {
				ret = min(ret, ri-le+dfs(le, ch)+dfs(ch, ri))
			}
		}
		if ret >= inf {
			ret = 0
		}
		if mem[le] == nil {
			mem[le] = make(map[int]int)
		}
		mem[le][ri] = ret
		return ret
	}
	return dfs(0, n)
}
