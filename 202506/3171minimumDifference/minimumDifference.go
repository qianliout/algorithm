package main

import (
	"math"
)

func main() {

}

func minimumDifference2(nums []int, k int) int {
	ans := math.MaxInt64

	for i, x := range nums {
		ans = min(ans, abs(x-k))
		j := i - 1
		for j >= 0 && nums[j]|x != nums[j] {
			nums[j] = nums[j] | x
			ans = min(ans, abs(nums[j]-k))
			j--
		}
	}
	return ans
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func minimumSubarrayLength(nums []int, k int) int {
	inf := math.MaxInt64 / 10
	ans := inf

	for i, x := range nums {
		// 值 至少 为 k
		if x >= k {
			return 1
		}
		j := i - 1
		for j >= 0 && nums[j]|x != nums[j] {
			nums[j] = nums[j] | x
			if nums[j] >= k {
				ans = min(ans, i-j+1)
			}
			j--
		}
	}
	if ans >= inf {
		return -1
	}
	return ans
}
