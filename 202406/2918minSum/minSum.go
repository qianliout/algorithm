package main

import (
	"slices"
)

func main() {

}

func minSum(nums1 []int, nums2 []int) int64 {
	c1 := slices.Contains(nums1, 0)
	c2 := slices.Contains(nums2, 0)
	s1, s2 := 0, 0
	for _, ch := range nums1 {
		s1 += max(1, ch)
	}
	for _, ch := range nums2 {
		s2 += max(1, ch)
	}
	// 情况一
	if !c1 && !c2 && s1 != s2 {
		return -1
	}
	// 情况二
	if !c1 && c2 && s1 < s2 {
		return -1
	}
	if c1 && !c2 && s2 < s1 {
		return -1
	}

	return int64(max(s1, s2))
}
