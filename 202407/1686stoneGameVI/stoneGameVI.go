package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(stoneGameVI([]int{1, 2}, []int{3, 1}))
}

func stoneGameVI(aliceValues []int, bobValues []int) int {
	n := len(aliceValues)
	pairs := make([]pair, n)
	for i := range aliceValues {
		pairs[i] = pair{aliceValues[i], bobValues[i], aliceValues[i] + bobValues[i]}
	}
	sort.Slice(pairs, func(i, j int) bool { return pairs[i].c >= pairs[j].c })

	diff := 0
	for i := 0; i < n; i++ {
		if i&1 == 0 {
			diff += pairs[i].a
		} else {
			diff -= pairs[i].b
		}
	}
	if diff > 0 {
		return 1
	}
	if diff < 0 {
		return -1
	}
	return 0
}

type pair struct {
	a, b, c int
}
