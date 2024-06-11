package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxNonDecreasingLength([]int{2, 3, 1}, []int{1, 2, 1}))
	fmt.Println(maxNonDecreasingLength([]int{11, 7, 7, 9}, []int{19, 19, 1, 7}))
}

func maxNonDecreasingLength(nums1 []int, nums2 []int) int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return 0
	}
	ans, n, cnt := 1, len(nums1), 1

	pre := min(nums1[0], nums2[0])

	for i := 1; i < n; i++ {
		if max(nums1[i], nums2[i]) < pre {
			ans = max(ans, cnt)
			cnt = 1
			pre = min(nums1[i], nums2[i])
			continue
		}
		if nums1[i] < pre {
			pre = nums2[i]
			cnt++
		} else if nums2[i] < pre {
			pre = nums1[i]
			cnt++
		} else {
			pre = min(nums1[i], nums2[i])
			cnt++
		}
	}
	ans = max(ans, cnt)
	return ans
}
