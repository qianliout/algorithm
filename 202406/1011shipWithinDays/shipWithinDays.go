package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(f([]int{1, 2, 3, 1, 1}, 2))
	// fmt.Println(f([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 100))
	fmt.Println(shipWithinDays([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5))
	fmt.Println(shipWithinDays([]int{1, 2, 3, 1, 1}, 4))
}

func shipWithinDays(w []int, d int) int {
	sum := 0

	for _, ch := range w {
		sum += ch
	}
	mi := slices.Max(w) // 最大的货物，这个船也得装下
	le, ri := mi, sum

	for le < ri {
		mid := le + (ri-le)/2
		// 左端点
		if mid >= mi && mid <= sum && f(w, mid) <= d {
			ri = mid
		} else {
			le = mid + 1
		}
	}
	if ri <= 0 || f(w, le) > d {
		return -1
	}
	return le
}

// 在船的能力是 t 时最少要多少天
func f(w []int, t int) int {
	ans := 0
	pre := 0
	for _, ch := range w {
		if pre+ch > t {
			ans++
			pre = ch
		} else {
			pre += ch
		}
	}
	if pre > 0 {
		ans++
	}
	return ans
}
