package main

import (
	"sort"
)

func main() {

}

func smallestRangeII(nums []int, k int) int {
	sort.Ints(nums)
	n := len(nums)
	ans := nums[n-1] - nums[0]
	for i := 0; i < n-1; i++ {
		mx := max(nums[i]+k, nums[n-1]-k)
		mi := min(nums[0]+k, nums[i+1]-k)
		ans = min(ans, abs(mx-mi))
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
