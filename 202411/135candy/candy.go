package main

import (
	"fmt"
)

func main() {
	fmt.Println(candy([]int{1, 0, 2}))
}

func candy(ratings []int) int {
	n := len(ratings)
	left := make([]int, n)
	right := make([]int, n)
	for i := 0; i < n; i++ {
		if i > 0 && ratings[i] > ratings[i-1] {
			left[i] = left[i-1] + 1
		} else {
			left[i] = 1
		}
	}
	for i := n - 1; i >= 0; i-- {
		if i < n-1 && ratings[i] > ratings[i+1] {
			right[i] = right[i+1] + 1
		} else {
			right[i] = 1
		}
	}

	ans := 0
	for i := 0; i < n; i++ {
		ans += max(left[i], right[i])
	}
	return ans
}
