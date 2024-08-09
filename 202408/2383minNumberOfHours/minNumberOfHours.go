package main

import (
	"fmt"
)

func main() {
	fmt.Println(minNumberOfHours(5, 3, []int{1, 4, 3, 2}, []int{2, 6, 3, 1}))
}

func minNumberOfHours(en int, ex int, energy []int, experience []int) int {
	// 依次击败对手，所以不能排序
	n := len(energy)
	ans := 0
	for i := 0; i < n; i++ {
		x, y := energy[i], experience[i]
		if en <= x {
			ans += x + 1 - en
			en = x + 1
		}
		en -= x
		if ex <= y {
			ans += y + 1 - ex
			ex = y + 1
		}
		ex += y
	}
	return ans
}
