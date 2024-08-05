package main

import (
	"sort"
)

func main() {

}

func countElements(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	le, ri := 0, n-1
	for le <= ri {
		if nums[le] == nums[0] {
			le++
			continue
		} else if nums[ri] == nums[n-1] {
			ri--
			continue
		}
		break
	}
	return max(0, ri-le+1)
}
