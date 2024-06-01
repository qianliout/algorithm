package main

import (
	"sort"
)

func main() {

}

func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	n := len(worker)
	type pair struct{ d, p, idx int }
	pairs := make([]pair, 0)
	for i := 0; i < n; i++ {
		pairs = append(pairs, pair{d: difficulty[i], p: profit[i], idx: i})
	}
	sort.Slice(pairs, func(i, j int) bool { return pairs[i].d < pairs[j].d })
	sort.Ints(worker)

	ans := 0
	maxP := 0
	j := 0
	for _, wo := range worker {
		for j < n && pairs[j].d <= wo {
			maxP = max(maxP, pairs[j].p)
			j++
		}
		ans += maxP
	}

	return ans
}
