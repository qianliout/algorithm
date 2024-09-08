package main

import (
	"sort"
)

func main() {

}

func minimumAverage(nums []int) float64 {
	sort.Ints(nums)
	n := len(nums)
	ans := nums[0] + nums[n-1]
	for i := 1; i < n/2; i++ {
		ans = min(ans, nums[i]+nums[n-i-1])
	}
	return float64(ans) / 2
}
