package main

import (
	"slices"
)

func main() {

}

func countSubarrays(nums []int, k int) int64 {
	n := len(nums)
	mx := slices.Max(nums)
	ans := 0
	le, ri := 0, 0
	cnt := 0
	for le <= ri && ri < n {
		if nums[ri] == mx {
			cnt++
		}
		for cnt == k {
			if nums[le] == mx {
				cnt--
			}
			le++
		}
		ans += le
		ri++
	}
	return int64(ans)
}
