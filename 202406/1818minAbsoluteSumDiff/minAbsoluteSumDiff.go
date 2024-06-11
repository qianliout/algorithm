package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(minAbsoluteSumDiff([]int{1, 7, 5}, []int{2, 3, 5}))
}

func minAbsoluteSumDiff(nums1 []int, nums2 []int) int {
	base := int(math.Pow10(9)) + 7
	n := len(nums1)
	rec := make(sort.IntSlice, n)
	for i := range rec {
		rec[i] = nums1[i]
	}
	rec.Sort()
	sum := 0
	sub := 0
	for i, x := range nums2 {
		diff := abs(x - nums1[i])
		sum += diff
		// 找最接近x的其他值，可能在 x的左边，也可能在右边
		j := rec.Search(x)
		if j < n {
			sub = max(sub, diff-abs(rec[j]-x))
		}
		if j > 0 {
			sub = max(sub, diff-abs(rec[j-1]-x))
		}
	}
	return (sum - sub) % base

}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
