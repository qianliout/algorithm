package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(deckRevealedIncreasing([]int{17, 13, 11, 2, 3, 5, 7}))
}

// 每个值都不一样
// 模拟题目
func deckRevealedIncreasing(deck []int) []int {
	tem := make([]int, 0)
	dq := make([]int, 0)
	tem = append(tem, deck...)
	dq = append(dq, deck...)
	n := len(deck)
	i := 0
	idx := make(map[int]int)
	for len(dq) > 0 {
		fir := dq[0]
		idx[fir] = i
		i++
		dq = dq[1:]
		if len(dq) > 0 {
			se := dq[0]
			dq = dq[1:]
			dq = append(dq, se)
		}
	}
	sort.Ints(tem)
	ans := make([]int, n)
	for j := 0; j < n; j++ {
		ans[j] = tem[idx[deck[j]]]
	}
	return ans
}
