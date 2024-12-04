package main

import (
	"fmt"
	"sort"
)

func main() {
	// fmt.Println(minCost([]int{4, 4, 4, 4, 3}, []int{5, 5, 5, 5, 3}))
	fmt.Println(minCost([]int{84, 80, 43, 8, 80, 88, 43, 14, 100, 88}, []int{32, 32, 42, 68, 68, 100, 42, 84, 14, 8}))
}

func minCost(basket1 []int, basket2 []int) int64 {
	if len(basket1) != len(basket2) {
		return -1
	}
	am, bm := make(map[int]int), make(map[int]int)
	all := make(map[int]int)

	n := len(basket1)
	for i := 0; i < n; i++ {
		am[basket1[i]]++
		bm[basket2[i]]++
		all[basket1[i]]++
		all[basket2[i]]++
	}
	for _, v := range all {
		if v&1 == 1 {
			return -1
		}
	}
	b1, b2 := make([]int, 0), make([]int, 0)
	for _, ch := range basket1 {
		if bm[ch] <= 0 {
			b1 = append(b1, ch)
		} else {
			bm[ch]--
		}
	}

	for _, ch := range basket2 {
		if am[ch] <= 0 {
			b2 = append(b2, ch)
		} else {
			am[ch]--
		}
	}
	sort.Ints(basket1)
	sort.Ints(basket2)
	sort.Ints(b1)
	sort.Slice(b2, func(i, j int) bool {
		return b2[i] > b2[j]
	})
	ans := 0
	cnt := len(b1)
	for i := 0; i < cnt; i += 2 {
		ans += min(b1[i], b2[i])
	}
	return int64(ans)
}
