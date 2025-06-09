package main

import (
	"fmt"
)

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}

func trap(height []int) int {
	n := len(height)
	lx, rm := height[0], height[n-1]
	ans := 0
	le, ri := 0, n-1
	for le < ri {
		lx = max(lx, height[le])
		rm = max(rm, height[ri])
		if lx < rm {
			ans += lx - height[le]
			le++
		} else {
			ans += rm - height[ri]
			ri--
		}
	}

	return ans
}
