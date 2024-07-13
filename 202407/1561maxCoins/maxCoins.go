package main

import (
	"sort"
)

func main() {

}

// 不用想那么复杂，先排序，分3份，第一份是最少的，不要！
// 剩下的跳着拿！
func maxCoins(piles []int) int {
	sort.Ints(piles)
	n := len(piles) / 3
	ans := 0
	for i := n; i < 3*n; i = i + 2 {
		ans += piles[i]
	}
	return ans
}
