package main

import (
	"fmt"
)

func main() {
	fmt.Println(subarrayLCM([]int{3, 6, 2, 7, 1}, 6))
}

func subarrayLCM(nums []int, k int) int {
	ans, n := 0, len(nums)
	for i := range nums {
		res := 1
		for j := i; j < n; j++ {
			res = res / gcb(res, nums[j]) * nums[j]
			if k%res > 0 {
				break
			}
			if res == k {
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

func lcm(a, b int) int {
	g := gcb(a, b)
	return a * b / g
}
