package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minCost([]int{1, 3, 5, 2}, []int{2, 3, 1, 14}))
}

func minCost(nums []int, cost []int) int64 {
	total := 0

	sumCost := 0
	for _, ch := range cost {
		sumCost += ch
	}
	n := len(nums)
	pairs := make([]pair, n)
	for i := range pairs {
		pairs[i] = pair{nums[i], cost[i]}
	}

	sort.Slice(pairs, func(i, j int) bool { return pairs[i].x < pairs[j].x })

	for _, ch := range pairs {
		total += ch.cost * (ch.x - pairs[0].x)
	}

	mi := total
	for i := 1; i < n; i++ {
		sumCost -= pairs[i-1].cost * 2
		total -= sumCost * (pairs[i].x - pairs[i-1].x)
		mi = min(mi, total)
	}
	return int64(mi)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

type pair struct {
	x, cost int
}
