package main

import (
	"sort"
)

func main() {

}

func numMovesStonesII(nums []int) []int {
	sort.Ints(nums)
	n := len(nums)
	if n <= 2 {
		return []int{0, 0}
	}
	e1 := nums[n-2] - nums[0] - (n - 2)
	e2 := nums[n-1] - nums[1] - (n - 2)
	maxMove := max(e2, e1)
	if e1 == 0 || e2 == 0 {
		return []int{min(2, maxMove), maxMove}
	}
	maxCnt, left := 0, 0
	for ri, sr := range nums {
		for sr-nums[left]+1 > n {
			left++
		}
		maxCnt = max(maxCnt, ri-left+1)
	}
	return []int{n - maxCnt, maxMove}
}
