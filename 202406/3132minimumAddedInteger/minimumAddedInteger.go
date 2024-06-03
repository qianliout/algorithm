package main

import (
	"fmt"
	"slices"
	"sort"
)

func main() {
	fmt.Println(minimumAddedInteger([]int{4, 20, 16, 12, 8}, []int{14, 18, 10}))
}

func minimumAddedInteger(nums1 []int, nums2 []int) int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	for i := 2; i >= 0; i-- {
		diff := nums2[0] - nums1[i]
		j := 0
		for k := i; k < len(nums1); k++ {
			if nums2[j] == nums1[k]+diff {
				j++
				if j >= len(nums2) {
					// 因为上面是到着枚举的，后面的结果只会更小，所以可以直接返回
					return diff
				}
			}
		}

	}
	return nums2[0] - nums1[0]
}

func addedInteger(nums1, nums2 []int) int {
	// return slices.Min(nums2) - slices.Min(nums1)
	return slices.Max(nums2) - slices.Max(nums1)
}
