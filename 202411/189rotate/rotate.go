package main

import (
	"fmt"
)

func main() {
	rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3)
}
func rotate2(nums []int, k int) {
	n := len(nums)
	k = k % n
	left := make([]int, 0)
	for i := n - k; i < n; i++ {
		left = append(left, nums[i])
	}

	j := n - 1
	for i := n - k - 1; i >= 0; i-- {
		nums[j] = nums[i]
		j--
	}
	for i := 0; i < len(left); i++ {
		nums[i] = left[i]
	}
	fmt.Println(nums)
}

func rotate(nums []int, k int) {
	k = k % len(nums)
	rotateAll(nums)
	rotateAll(nums[:k])
	rotateAll(nums[k:])
}

func rotateAll(nums []int) {
	l, r := 0, len(nums)-1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}
