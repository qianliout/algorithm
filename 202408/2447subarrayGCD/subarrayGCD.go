package main

import (
	"fmt"
)

func main() {
	fmt.Println(subarrayGCD([]int{9, 3, 1, 2, 6, 3}, 3))
}

func subarrayGCD(nums []int, k int) int {
	ans := 0
	n := len(nums)
	for i := range nums {
		g := 0
		for j := i; j < n; j++ {
			g = gcb(g, nums[j])
			if g%k > 0 {
				break
			}
			if g == k {
				ans++
			}
		}
	}
	return ans
}

func gcb(a, b int) int {
	if b == 0 {
		return a
	}
	return gcb(b, a%b)
}
