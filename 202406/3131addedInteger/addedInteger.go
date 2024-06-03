package main

import (
	"slices"
)

func main() {

}

func addedInteger(nums1, nums2 []int) int {
	// return slices.Min(nums2) - slices.Min(nums1)
	return slices.Max(nums2) - slices.Max(nums1)
}
