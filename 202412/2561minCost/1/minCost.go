package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	// fmt.Println(minCost([]int{4, 4, 4, 4, 3}, []int{5, 5, 5, 5, 3}))
	fmt.Println(minCost([]int{84, 80, 43, 8, 80, 88, 43, 14, 100, 88}, []int{32, 32, 42, 68, 68, 100, 42, 84, 14, 8}))
}

func minCost(basket1 []int, basket2 []int) int64 {
	cnt := make(map[int]int)
	if len(basket1) != len(basket2) {
		return -1
	}
	n := len(basket1)
	for i := 0; i < n; i++ {
		cnt[basket1[i]]++
		cnt[basket2[i]]--
	}
	mi, find := math.MaxInt, make([]int, 0)
	for x, c := range cnt {
		mi = min(mi, x)
		if abs(c)%2 != 0 {
			return -1
		}
		for c = abs(c) / 2; c > 0; c-- {
			find = append(find, x)
		}
	}
	// mi 是一个 工具人，可以用这个工具人去交换两次
	ans := 0
	sort.Ints(find)
	for i := 0; i < len(find)/2; i++ {
		ans += min(find[i], mi*2)
	}
	return int64(ans)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
