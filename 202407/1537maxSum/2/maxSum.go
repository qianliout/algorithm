package main

import (
	"math"
)

func main() {
}

// 你有两个 有序 且数组内元素互不相同的数组 nums1 和 nums2 。
func maxSum2(nums1 []int, nums2 []int) int {
	i, j, sum1, sum2, mod := 0, 0, 0, 0, int(math.Pow10(9))+7
	n1, n2 := len(nums1), len(nums2)
	for i < n1 && j < n2 {
		if nums1[i] < nums2[j] {
			sum1 = (sum1 + nums1[i]) % mod // 不能提前取余数，不知道为啥
			i++
		} else if nums1[i] > nums2[j] {
			sum2 = (sum2 + nums2[j]) % mod // 不能提前取余数，不知道为啥
			j++
		} else {
			a := (max(sum1, sum2) + nums1[i]) % mod
			sum1, sum2 = a, a
			i++
			j++
		}
	}
	// 把剩余的加上
	for ; i < n1; i++ {
		sum1 = (sum1 + nums1[i]) % mod
	}
	for ; j < n2; j++ {
		sum2 = (sum2 + nums2[j]) % mod
	}
	return max(sum1, sum2) % mod
}

// 你有两个 有序 且数组内元素互不相同的数组 nums1 和 nums2 。
func maxSum(nums1 []int, nums2 []int) int {
	i, j, sum1, sum2, mod := 0, 0, 0, 0, int(math.Pow10(9))+7
	n1, n2 := len(nums1), len(nums2)
	for i < n1 && j < n2 {
		if nums1[i] < nums2[j] {
			sum1 += nums1[i]
			i++
		} else if nums1[i] > nums2[j] {
			sum2 += nums2[j]
			j++
		} else {
			a := (max(sum1, sum2) + nums1[i]) % mod
			sum1, sum2 = a, a
			i++
			j++
		}
	}
	// 把剩余的加上
	for ; i < n1; i++ {
		sum1 += nums1[i]
	}
	for ; j < n2; j++ {
		sum2 += nums2[j]
	}
	return max(sum1, sum2) % mod
}
