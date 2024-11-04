package main

import (
	"sort"
)

func main() {

}

func numberGame(nums []int) []int {
	sort.Ints(nums)
	n := len(nums)
	for i := 0; i < n; i += 2 {
		nums[i+1], nums[i] = nums[i], nums[i+1]
	}
	return nums
}
