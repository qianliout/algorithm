package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minPairSum([]int{3, 5, 2, 3}))
}

func minPairSum(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	ans := nums[0] + nums[n-1-0]
	for i := 0; i < n/2; i++ {
		ans = max(ans, nums[i]+nums[n-1-i])
	}
	return ans
}
