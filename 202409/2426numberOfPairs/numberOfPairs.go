package main

import (
	"sort"
)

func main() {

}

// 暴力做会会超时
func numberOfPairs1(nums1 []int, nums2 []int, diff int) int64 {
	n := len(nums1)
	ans := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if (nums1[i]-nums2[i])-(nums1[j]-nums2[j]) <= diff {
				ans++
			}
		}
	}
	return int64(ans)
}

func numberOfPairs(nums1 []int, nums2 []int, diff int) int64 {
	n := len(nums1)
	if n == 0 || n == 1 {
		return 0 // 处理边界情况
	}

	diffArr := make([]int, n)
	for i := 0; i < n; i++ {
		diffArr[i] = nums1[i] - nums2[i]
	}

	sort.Ints(diffArr) // 对差值数组进行排序

	ans := 0
	for i := 0; i < n; i++ {
		left, right := i, n-1
		for left < right {
			mid := (left + right + 1) / 2
			if diffArr[mid]-diffArr[i] > diff {
				right = mid - 1
			} else {
				left = mid
			}
		}
		ans += left - i
	}
	return int64(ans)
}
