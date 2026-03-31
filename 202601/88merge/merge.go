package main

import (
	"sort"
)

func main() {}

func merge(nums1 []int, m int, nums2 []int, n int) {
	i := m
	for j := 0; i < n; j++ {
		nums1[i] = nums2[j]
		i++
	}
	sort.Ints(nums1)
}
