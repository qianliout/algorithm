package main

import (
	"sort"
)

func main() {

}

func kIncreasing(arr []int, k int) int {
	n := len(arr)
	ans := 0
	for i := 0; i < k && i < n; i++ {
		f := make([]int, 0)
		for j := i; j < n; j += k {
			ch := arr[j]
			p := sort.SearchInts(f, ch+1)
			if p < len(f) {
				f[p] = ch
			} else {
				f = append(f, ch)
			}
		}
		ans += len(f)
	}
	return n - ans
}

// 最大递增子序列
// 不是严格递增
func lengthOfLIS1(nums []int) int {
	n := len(nums)
	dp := make([]int, 0)
	for i := 0; i < n; i++ {
		p := sort.SearchInts(dp, nums[i]+1)
		if p < len(dp) {
			dp[p] = nums[i]
		} else {
			dp = append(dp, nums[i])
		}
	}
	return n - len(dp)
}
