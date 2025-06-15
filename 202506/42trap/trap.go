package main

import (
	"fmt"
)

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})) // 6
}

func trap1(height []int) int {
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

func trap(height []int) int {
	ans := 0
	st := make([]int, 0)
	for i, c := range height {
		for len(st) > 0 && c > height[st[len(st)-1]] {
			bottom := height[st[len(st)-1]]
			st = st[:len(st)-1]
			if len(st) == 0 {
				break
			}
			left := st[len(st)-1]
			h := min(height[left], c) - bottom
			ans += h * (i - left - 1)
		}
		st = append(st, i)
	}
	return ans
}
