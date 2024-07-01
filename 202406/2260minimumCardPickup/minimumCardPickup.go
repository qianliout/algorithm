package main

import (
	"fmt"
)

func main() {
	fmt.Println(minimumCardPickup([]int{1, 0, 5, 3}))
}

func minimumCardPickup(cards []int) int {
	n := len(cards)
	wind := make(map[int]int)
	ans := n + 1
	le, ri := 0, 0
	for le <= ri && ri < n {
		num := cards[ri]
		wind[num]++
		for le <= ri && wind[num] >= 2 {
			ans = min(ans, ri-le+1)
			wind[cards[le]]--
			le++
		}
		ri++
	}
	if ans > n {
		return -1
	}
	return ans
}
