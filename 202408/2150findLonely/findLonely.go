package main

import (
	"sort"
)

func main() {

}

func findLonely(nums []int) []int {
	ans, n := make([]int, 0), len(nums)
	sort.Ints(nums)
	for i := 0; i < n; i++ {
		if i > 0 && abs(nums[i]-nums[i-1]) <= 1 {
			continue
		}
		if i < n-1 && abs(nums[i]-nums[i+1]) <= 1 {
			continue
		}
		ans = append(ans, nums[i])
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
