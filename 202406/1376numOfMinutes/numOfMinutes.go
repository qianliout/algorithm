package main

import (
	"fmt"
)

func main() {
	fmt.Println(numOfMinutes(11, 4, []int{5, 9, 6, 10, -1, 8, 9, 1, 9, 3, 4}, []int{0, 213, 0, 253, 686, 170, 975, 0, 261, 309, 337}))
}

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	g := make([][]int, n)
	for i, x := range manager {
		if x != -1 {
			g[x] = append(g[x], i)
		}
	}
	ans := 0
	queue := make([]int, 0)
	queue = append(queue, headID)

	for len(queue) > 0 {
		ti := 0
		lev := make([]int, 0)
		for _, no := range queue {
			ti = max(ti, informTime[no])
			for _, nx := range g[no] {
				lev = append(lev, nx)
			}
		}
		ans += ti
		queue = lev
	}
	return ans
}
