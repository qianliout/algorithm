package main

import (
	"sort"
)

func main() {

}

func reductionOperations(nums []int) int {
	sort.Ints(nums)
	op := 0
	ans := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			op++
		}
		ans += op
	}
	return ans
}
