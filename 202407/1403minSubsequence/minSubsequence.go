package main

import (
	"sort"
)

func main() {

}

func minSubsequence(nums []int) []int {
	sort.Ints(nums)
	n := len(nums)
	sum := 0
	for _, ch := range nums {
		sum += ch
	}
	ans := make([]int, 0)
	cnt := 0
	for i := n - 1; i >= 0; i-- {
		ans = append(ans, nums[i])
		cnt += nums[i]
		if cnt > (sum - cnt) {
			break
		}
	}
	return ans
}
