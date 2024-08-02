package main

import (
	"sort"
)

func main() {

}

func targetIndices(nums []int, target int) []int {
	n := len(nums)
	pairs := make([]pair, n)
	for i, ch := range nums {
		pairs[i] = pair{i, ch}
	}
	sort.SliceStable(pairs, func(i, j int) bool { return pairs[i].val <= pairs[j].val })

	ans := make([]int, 0)
	for i := 0; i < n; i++ {
		if pairs[i].val == target {
			ans = append(ans, i)
		}
	}
	return ans
}

type pair struct {
	idx int
	val int
}
