package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(5 / 5)
}

func minLengthAfterRemovals(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	x := nums[n/2]
	ri := sort.SearchInts(nums, x+1)
	le := sort.SearchInts(nums, x)
	maxCnt := ri - le
	return max(maxCnt*2-n, n%2)
}
