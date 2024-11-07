package main

import (
	"fmt"
)

func main() {
	fmt.Println(minSubArrayLen(9, []int{1, 1, 1, 1, 1, 1, 1, 1}))
}

// 能得到正确结果，但是会超时
func minSubArrayLen1(k int, nums []int) int {
	n := len(nums)
	pre := make([]int, n+1)
	for i, ch := range nums {
		pre[i+1] = pre[i] + ch
	}
	ans := n + 1
	for i := 1; i <= n; i++ {
		for j := i - 1; j >= 0; j-- {
			if pre[i]-pre[j] >= k {
				ans = min(ans, i-j)
			}
		}
	}
	if ans == n+1 {
		return 0
	}
	return ans
}

func minSubArrayLen(k int, nums []int) int {
	n := len(nums)
	sum := make([]int, n+1)
	for i, ch := range nums {
		sum[i+1] = ch + sum[i]
	}
	ans := n + 1
	left, right := 0, 1
	for left <= right && right <= n {
		for left <= right && sum[right]-sum[left] >= k {
			ans = min(ans, right-left)
			left++
		}
		right++
	}

	if ans == n+1 {
		return 0
	}
	return ans
}
