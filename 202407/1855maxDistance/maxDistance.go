package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxDistance([]int{9819, 9508, 7398, 7347, 6337, 5756, 5493, 5446, 5123, 3215, 1597, 774, 368, 313}, []int{9933, 9813, 9770, 9697, 9514, 9490, 9441, 9439, 8939, 8754, 8665, 8560}))
}

func maxDistance(nums1 []int, nums2 []int) int {
	// 根据题目中的意思，nums1和nums2的意义不一样
	// if len(nums1) > len(nums2) {
	// 	nums1, nums2 = nums2, nums1
	// }
	n := len(nums2)
	ans := 0
	for i := 0; i < len(nums1); i++ {
		le, ri := 0, n
		for le < ri {
			mid := le + (ri-le+1)/2
			if mid >= 0 && mid < n && nums1[i] <= nums2[mid] {
				le = mid
			} else {
				ri = mid - 1
			}
		}
		if le < n && i <= le {
			ans = max(ans, le-i)
		}
	}

	return ans
}
