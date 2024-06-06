package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minimumSum([]int{5, 4, 8, 7, 10, 2}))
}

func minimumSum(nums []int) int {
	ans := math.MaxInt
	n := len(nums)
	if n < 3 {
		return 0
	}
	suf := make([]int, n)

	suf[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		suf[i] = min(suf[i+1], nums[i])
	}
	pre := nums[0]
	for i := 1; i < n-1; i++ {
		if pre < nums[i] && nums[i] > suf[i+1] {
			ans = min(ans, pre+nums[i]+suf[i+1])
		}
		pre = min(pre, nums[i])
	}
	if ans == math.MaxInt {
		return -1
	}
	return ans
}
