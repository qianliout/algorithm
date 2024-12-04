package main

import (
	"math"
	"sort"
)

func main() {

}

func minCost(basket1, basket2 []int) (ans int64) {
	cnt := map[int]int{}
	for i, x := range basket1 {
		cnt[x]++
		cnt[basket2[i]]--
	}
	mn, a := math.MaxInt, []int{}
	for x, c := range cnt {
		if c%2 != 0 {
			return -1
		}
		mn = min(mn, x)
		for c = abs(c) / 2; c > 0; c-- {
			a = append(a, x)
		}
	}
	sort.Ints(a) // 也可以用快速选择
	for _, x := range a[:len(a)/2] {
		ans += int64(min(x, mn*2))
	}
	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
