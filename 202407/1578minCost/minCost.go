package main

import (
	"fmt"
)

func main() {
	fmt.Println(minCost("abaac", []int{1, 2, 3, 4, 5}))
	fmt.Println(minCost("abc", []int{1, 2, 3}))
	fmt.Println(minCost("aabaa", []int{1, 2, 3, 4, 1}))
}

func minCost(colors string, neededTime []int) int {
	ans := 0
	st := make([]int, 0)
	n := len(neededTime)
	le := 0
	for le < n {
		add := true
		for len(st) > 0 && colors[st[len(st)-1]] == colors[le] {
			last := neededTime[st[len(st)-1]]
			cur := neededTime[le]
			if last >= cur {
				add = false
				ans += cur
				break
			} else {
				ans += last
				st = st[:len(st)-1]
			}
		}
		if add {
			st = append(st, le)
		}
		le++
	}
	return ans
}
