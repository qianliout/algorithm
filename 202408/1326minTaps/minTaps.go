package main

import (
	"fmt"
)

func main() {
	fmt.Println(minTaps(5, []int{4, 1, 1, 1, 1, 1}))
	fmt.Println(minTaps(5, []int{3, 4, 1, 1, 0, 0}))
}

func minTaps(n int, ranges []int) int {
	rightMost := make([]int, n+1)
	for i, ch := range ranges {
		left := max(0, i-ch)
		rightMost[left] = max(rightMost[left], i+ch)
	}
	current := 0
	next := 0
	ans := 0
	// 这一步，i<n 不容易理解，还没有理解
	for i := 0; i < n; i++ {
		next = max(next, rightMost[i])
		if i == current {
			// 假如 i=n-1,此时 current=next=n-1,此时也会返回-1,所以上面i<n 是对的
			if current == next {
				return -1
			}
			ans++
			current = next
		}
	}
	return ans
}
