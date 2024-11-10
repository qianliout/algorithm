package main

import (
	"sort"
)

func main() {

}

func minOperationsToMakeMedianK1(nums []int, k int) int64 {
	sort.Ints(nums)
	n := len(nums)
	mid := n / 2
	ans := 0
	if nums[mid] > k {
		for i := mid; i >= 0; i-- {
			ans += max(nums[i]-k, 0)
		}
	} else {
		for i := mid; i < n; i++ {
			ans += max(k-nums[i], 0)
		}
	}
	return int64(ans)
}

func minOperationsToMakeMedianK(nums []int, k int) int64 {
	sort.Ints(nums)
	n := len(nums)
	mid := n / 2
	ans := 0
	if nums[mid] > k {

		for i := mid; i >= 0; i-- {
			if nums[i] < k {
				break
			}
			ans += nums[i] - k
		}

	} else {
		for i := mid; i < n; i++ {
			if nums[i] > k {
				break
			}
			ans += k - nums[i]
		}
	}
	return int64(ans)
}
