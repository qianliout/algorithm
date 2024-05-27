package main

import (
	"sort"
)

func main() {

}

func minIncrementForUnique(nums []int) int {
	sort.Ints(nums)
	if len(nums) <= 1 {
		return 0
	}
	pre := nums[0]
	res := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] == pre+1 {
			pre = nums[i]
		} else if nums[i] > pre+1 {
			pre = nums[i]
		} else {
			res += pre - nums[i] + 1
			pre = pre + 1
		}
	}
	return res
}
