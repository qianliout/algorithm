package main

import (
	"fmt"
)

func main() {
	// fmt.Println(hIndex([]int{0, 1, 3, 5, 6}))
	fmt.Println(hIndex([]int{0, 1}))
}

func hIndex(citations []int) int {
	n := len(citations)
	le, ri := 0, n
	for le < ri {
		mid := le + (ri-le)/2
		if mid >= 0 && mid < n && citations[mid] >= n-mid {
			ri = mid
		} else {
			le = mid + 1
		}
	}
	if le < 0 || le >= n || citations[le] < n-le {
		return 0
	}
	return n - le
}

// h 代表“高引用次数”（high citations），一名科研人员的 h 指数是指他（她）的 （n 篇论文中）至少 有 h 篇论文分别被引用了至少 h 次。
