package main

import (
	"slices"
)

func main() {

}
func smallestRangeI(nums []int, k int) int {
	mx := slices.Max(nums)
	mi := slices.Min(nums)

	if 2*k >= mx-mi {
		return 0
	}
	return mx - mi - 2*k
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
