package main

import (
	"fmt"
)

func main() {
	fmt.Println(minTaps(5, []int{3, 4, 1, 1, 0, 0}))
	fmt.Println(minTaps(5, []int{4, 1, 1, 1, 1, 1}))
}

func minTaps(n int, ranges []int) int {
	rightMost := make([]int, n+1)
	for i, ch := range ranges {
		left := max(0, i-ch)
		rightMost[left] = max(rightMost[left], i+ch)
	}
	curRight, nextRight := 0, 0
	ans := 0
	// 注意这里是i<n，而不是i<=n，因为 i 表示的是 left
	for i := 0; i < n; i++ {
		nextRight = max(nextRight, rightMost[i])
		if curRight == i {
			if nextRight == i {
				return -1
			}
			curRight = nextRight
			ans++
		}
	}
	return ans
}
