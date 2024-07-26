package main

import (
	"fmt"
)

func main() {
	fmt.Println(addRungs([]int{1, 3, 5, 10}, 2))
	fmt.Println(addRungs([]int{3, 6, 8, 10}, 3))
}

// rungs = [1,3,5,10], dist = 2
func addRungs2(rungs []int, dist int) int {
	start, mx := 0, dist
	n := len(rungs)
	ans := 0
	for start < n {
		ch := rungs[start]
		// 这样会超时
		if mx < ch {
			ans++
			mx += dist
		} else {
			mx = ch + dist
			start++
		}
	}

	return ans
}

func addRungs(rungs []int, dist int) int {
	start, mx := 0, dist
	n := len(rungs)
	ans := 0
	for start < n {
		ch := rungs[start]
		// 这样会超时
		if mx < ch {
			a := (ch - mx + dist - 1) / dist
			mx += dist * a
			ans += a
		} else {
			mx = ch + dist
			start++
		}
	}

	return ans
}
