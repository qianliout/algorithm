package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minimumCost([]int{1, 2, 3}))
}

func minimumCost(cost []int) int {
	sort.Ints(cost)
	ans := 0
	n := len(cost)
	i, buy := n-1, 0
	for i >= 0 {
		if buy == 2 {
			i--
			buy = 0
			continue
		}
		ans += cost[i]
		buy++
		i--
	}
	return ans
}
