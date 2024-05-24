package main

import (
	"fmt"
)

func main() {
	fmt.Println(canSeePersonsCount([]int{10, 6, 8, 5, 11, 9}))
}

func canSeePersonsCount(heights []int) []int {
	ans := make([]int, len(heights))
	st := make([]int, 0)
	for i := len(heights) - 1; i >= 0; i-- {
		for len(st) > 0 && heights[i] > st[len(st)-1] {
			ans[i]++
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			ans[i]++
		}
		st = append(st, heights[i])
	}
	return ans
}
