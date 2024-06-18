package main

import (
	"fmt"
)

func main() {
	fmt.Println(partitionDisjoint([]int{1, 1, 1, 0, 6, 12}))
	fmt.Println(partitionDisjoint([]int{1, 1}))
	fmt.Println(partitionDisjoint([]int{26, 51, 40, 58, 42, 76, 30, 48, 79, 91}))
}

func partitionDisjoint(nums []int) int {
	n := len(nums)
	mx, mi := make([]int, n), make([]int, n)
	mx[0], mi[n-1] = nums[0], nums[n-1]
	for i := 1; i < n; i++ {
		mx[i] = max(mx[i-1], nums[i])
	}
	for i := n - 2; i >= 0; i-- {
		mi[i] = min(mi[i+1], nums[i])
	}
	for i := 0; i < n-1; i++ {
		if mx[i] <= mi[i+1] {
			return i + 1
		}
	}
	return 0
}
